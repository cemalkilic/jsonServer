package database

import (
    "database/sql"
    "fmt"
    "log"
    "os"
)

func NewMySQLDBHandler() *sql.DB {

    mysqlUsername := os.Getenv("MYSQL_USER")
    mysqlPassword := os.Getenv("MYSQL_PASS")
    mysqlDBName := os.Getenv("MYSQL_DB")
    mysqlPort := os.Getenv("MYSQL_PORT")
    mysqlHost := os.Getenv("MYSQL_HOST")

    connStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", mysqlUsername, mysqlPassword, mysqlHost, mysqlPort, mysqlDBName)

    database, err := sql.Open("mysql", connStr)
    if err != nil {
        log.Fatal(err)
    }

    return database
}
