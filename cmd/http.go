package cmd

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var httpCmd = &cobra.Command{
	Use:   "http",
	Short: "start the back order service gateway",
	Run:   runServeHTTPCmd,
}

func init() {
	serveCmd.AddCommand(httpCmd)
}

func runServeHTTPCmd(cmd *cobra.Command, args []string) {
	logger := log.Default()
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	d := initDB()
	mysqlDsn := d.ToDSN()
	orm, err := gorm.Open(mysql.Open(mysqlDsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	maxOpenConnections := viper.GetInt(MySQLMaxOpenConnections)
	maxIdleConnections := viper.GetInt(MySQLMaxIdleConnections)

	sqlDB, _ := orm.DB()
	sqlDB.SetMaxOpenConns(maxOpenConnections)
	sqlDB.SetMaxIdleConns(maxIdleConnections)
	sqlDB.SetConnMaxLifetime(200 * time.Minute)

	go func() {
		router := mux.NewRouter().PathPrefix("/auth/").Subrouter()
		router.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {}).Methods("POST")

		httpMux := http.NewServeMux()
		httpMux.Handle("/auth/", router)

		httpHandler := cors.AllowAll().Handler(httpMux)

		srv := &http.Server{
			Addr:         ":8080",
			Handler:      httpHandler,
			IdleTimeout:  60 * time.Second,
			ReadTimeout:  15 * time.Second,
			WriteTimeout: 15 * time.Second,
		}
		logger.Print("server started")
		err = srv.ListenAndServe()
		if err != nil {
			panic(err)
		}
	}()
	<-c
	logger.Print("server graceful shutdown")
}
