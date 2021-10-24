package mysql

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"

	"github.com/spf13/viper"
)

var database *sql.DB
var config = viper.New()

func init() {
	config.SetConfigFile(`config.json`)
	err := config.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func SetupConnection() {
	dbHost := config.GetString(`mysql.host`)
	dbPort := config.GetString(`mysql.port`)
	dbUser := config.GetString(`mysql.user`)
	dbPass := config.GetString(`mysql.pass`)
	dbName := config.GetString(`mysql.name`)
	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)

	dbConnection, err := sql.Open(`mysql`, connection)
	if err != nil {
		panic(err.Error())
	}
	err = dbConnection.Ping()
	if err != nil {
		log.Fatal(err)
	}
	database = dbConnection
}

//GetConnection ...
func GetConnection() *sql.DB {
	return database
}

func CloseConnection() {
	database.Close()
}
