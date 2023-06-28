package database

import (
	"database/sql"
	"fmt"
	"log"
	"securigroup/tech-test/utils"

	_ "github.com/denisenkom/go-mssqldb"
)

var db *sql.DB


// Creates a connection pool to the MSSQL db
func CreateConnection() (*sql.DB, error) {
    var server = utils.GetEnvVar("DB_SERVER")
    var user = utils.GetEnvVar("DB_USER")
    var password = utils.GetEnvVar("DB_PASSWORD")
    var database = utils.GetEnvVar("DB_DATABASE")
    var port = utils.GetEnvVar("DB_PORT")

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
