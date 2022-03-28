package cmd

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/grpc"

	"github.com/TranTheTuan/authen-go/app/domain/usecase"
	"github.com/TranTheTuan/authen-go/app/infrastructure/casbin"
	internalGrpc "github.com/TranTheTuan/authen-go/app/infrastructure/grpc"
	pbAuth "github.com/TranTheTuan/pbtypes/build/go/auth"
)

var grpcCmd = &cobra.Command{
	Use:   "grpc",
	Short: "start grpc server",
	Run:   runServeGRPCCmd,
}

func init() {
	serveCmd.AddCommand(grpcCmd)
}

func runServeGRPCCmd(cmd *cobra.Command, args []string) {
	logger := log.Default()
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	d := initDB()
	mysqlDsn := d.ToDSN()
	orm, err := gorm.Open("mysql", mysqlDsn)
	if err != nil {
		panic(err)
	}

	maxOpenConnections := viper.GetInt(MySQLMaxOpenConnections)
	maxIdleConnections := viper.GetInt(MySQLMaxIdleConnections)

	orm.DB().SetMaxOpenConns(maxOpenConnections)
	orm.DB().SetMaxIdleConns(maxIdleConnections)
	orm.DB().SetConnMaxLifetime(200 * time.Minute)

	go func() {
		casbin.InitFromSQLLite(orm, viper.GetString(RBACFilePath))

		authorUsecase := usecase.NewAuthorUsecase()
		authorizeServiceServer := internalGrpc.NewAuthorizeServiceServer(authorUsecase)

		s := grpc.NewServer()
		pbAuth.RegisterAuthorizeServiceServer(s, authorizeServiceServer)

		lis, err := net.Listen("tcp", ":8080")
		if err != nil {
			logger.Fatalln("Failed to listen:", err)
		}
		logger.Println("Serving gRPC on 0.0.0.0:8080")
		logger.Fatalln(s.Serve(lis))
	}()
	<-c
	logger.Print("server graceful shutdown")
}
