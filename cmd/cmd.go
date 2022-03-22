/*
 * Copyright (c) 2020. SmartOSC Solution team - All Rights Reserved
 * Unauthorized copying of this file, via any medium is strictly prohibited
 * Proprietary and confidential
 */

package cmd

import (
	"fmt"

	"github.com/spf13/viper"
)

type DBConfiguration struct {
	Username  string
	Password  string
	Database  string
	Host      string
	Port      string
	Loc       string
	Charset   string
	ParseTime string
}

func NewDBConfiguration(username string, password string, database string, host string, port string, loc string, charset string) *DBConfiguration {
	d := &DBConfiguration{Username: username, Password: password, Database: database, Host: host, Port: port, Loc: loc, Charset: charset}
	d.ParseTime = "True"
	return d
}

func initDB() *DBConfiguration {
	u := viper.GetString(MySQLUserName)
	p := viper.GetString(MySQLPassword)
	d := viper.GetString(MySQLDatabase)
	h := viper.GetString(MySQLHost)
	po := viper.GetString(MySQLPort)
	char := viper.GetString(MySQLCharset)
	l := viper.GetString(MySQLLoc)
	return NewDBConfiguration(u, p, d, h, po, l, char)
}

// ToDSN returns the mysql data source name based on configuration.
func (d *DBConfiguration) ToDSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%s&loc=%s", d.Username, d.Password, d.Host, d.Port, d.Database, d.Charset, d.ParseTime, d.Loc)
}
