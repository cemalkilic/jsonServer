package database

import (
    "database/sql"
    "log"

    _ "github.com/mattn/go-sqlite3"
)

var db *databaseStore

func Init() {

    sqliteDatabase, err := sql.Open("sqlite3", "./sqlite-database.db")
    if err != nil {
        log.Fatal(err)
    }

    db = NewDatabaseStore(sqliteDatabase)
}

func GetDB() *databaseStore {
    return db
}
