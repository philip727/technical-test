package database

import (
	"database/sql"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	"log"
)

var db *sql.DB

var server = "127.0.0.1"
var user = "sa"
var password = "YourPassword123"
var database = "SecuriGroup"
var port = "1433"

func CreateConnection() (*sql.DB, error) {
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;database=%s;port=%s", server, user, password, database, port)

	conn, err := sql.Open("mssql", connString)
	if err != nil {
		log.Fatal("Error opening database connection: ", err.Error())
	}

    err = conn.Ping()
	if err != nil {
        conn.Close()
		log.Fatal("Error establishing connection to the database", err.Error())
	}

    return conn, nil
}
