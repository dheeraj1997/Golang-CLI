package mysqldb

import (
	"fmt"
	"net/url"
	"time"

	"github.com/jmoiron/sqlx"

	// Sql driver
	_ "github.com/go-sql-driver/mysql"
)

const (
	defaultMysqlDBConnection = 20
	dbHost = "localhost"
	dbPort = "3306"
	dbUser = "testuser"
	dbPass = "password"
	dbName = "testDb"
)

// GetSQLDbConnection Get SQL Db connection
func GetSQLDbConnection() (*sqlx.DB, error) {
	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	val := url.Values{}
	val.Add("charset", "utf8")
	val.Add("parseTime", "true")
	val.Add("loc", "Local")
	dsn := fmt.Sprintf("%s?%s", connection, val.Encode())
	//fmt.Println("Connect db uri: " + dsn)
	dbConn, err := sqlx.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	err = dbConn.Ping()
	if err != nil {
		return nil, err
	}
	// https://github.com/go-sql-driver/mysql#important-settings
	dbConn.SetMaxOpenConns(defaultMysqlDBConnection)
	dbConn.SetMaxIdleConns(defaultMysqlDBConnection)
	dbConn.SetConnMaxLifetime(time.Minute * 3)
	return dbConn.Unsafe(), nil
}
