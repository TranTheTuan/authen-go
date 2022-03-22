package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	MySQLUserName           = "system-mysql-username"
	MySQLPassword           = "system-mysql-password"
	MySQLDatabase           = "system-mysql-database"
	MySQLHost               = "system-mysql-host"
	MySQLPort               = "system-mysql-port"
	MySQLCharset            = "system-mysql-charset"
	MySQLLoc                = "system-mysql-loc"
	MySQLMaxOpenConnections = "system-mysql-max-open-conns"
	MySQLMaxIdleConnections = "system-mysql-max-idle-conns"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "authen-go",
	Short: "authentication root cmd",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.bm.yaml)")
	initConfiguration()

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		viper.AddConfigPath(home)
		viper.SetConfigName(".authen")
		viper.SetEnvPrefix("authen")
	}

	viper.SetEnvPrefix("authen")
	replacer := strings.NewReplacer("-", "_")
	viper.SetEnvKeyReplacer(replacer)
	viper.AutomaticEnv()

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

func initConfiguration() {
	rootCmd.PersistentFlags().String(MySQLUserName, "root", "mysql username")
	rootCmd.PersistentFlags().String(MySQLPassword, "admin123", "mysql password")
	rootCmd.PersistentFlags().String(MySQLDatabase, "polar_bear", "mysql database name")
	rootCmd.PersistentFlags().String(MySQLHost, "127.0.0.1", "mysql host")
	rootCmd.PersistentFlags().String(MySQLPort, "3306", "mysql port")
	rootCmd.PersistentFlags().String(MySQLCharset, "utf8mb4", "mysql default database character set. Recommend utf8mb4 for better Unicode support")
	rootCmd.PersistentFlags().String(MySQLLoc, "Local", "mysql password")
	rootCmd.PersistentFlags().String(MySQLMaxOpenConnections, "20", "mysql SetMaxOpenConns")
	rootCmd.PersistentFlags().String(MySQLMaxIdleConnections, "2", "mysql SetMaxIdleConns")

	//Bind flags to viper
	_ = viper.BindPFlag("system-mode", rootCmd.PersistentFlags().Lookup("system-mode"))
	_ = viper.BindPFlag("system-gorm-log-mode", rootCmd.PersistentFlags().Lookup("system-gorm-log-mode"))

	_ = viper.BindPFlag("system-mysql-username", rootCmd.PersistentFlags().Lookup("system-mysql-username"))
	_ = viper.BindPFlag("system-mysql-password", rootCmd.PersistentFlags().Lookup("system-mysql-password"))
	_ = viper.BindPFlag("system-mysql-database", rootCmd.PersistentFlags().Lookup("system-mysql-database"))
	_ = viper.BindPFlag("system-mysql-host", rootCmd.PersistentFlags().Lookup("system-mysql-host"))
	_ = viper.BindPFlag("system-mysql-port", rootCmd.PersistentFlags().Lookup("system-mysql-port"))
	_ = viper.BindPFlag("system-mysql-charset", rootCmd.PersistentFlags().Lookup("system-mysql-charset"))
	_ = viper.BindPFlag("system-mysql-loc", rootCmd.PersistentFlags().Lookup("system-mysql-loc"))
	_ = viper.BindPFlag(MySQLMaxOpenConnections, rootCmd.PersistentFlags().Lookup(MySQLMaxOpenConnections))
	_ = viper.BindPFlag(MySQLMaxIdleConnections, rootCmd.PersistentFlags().Lookup(MySQLMaxIdleConnections))
}
