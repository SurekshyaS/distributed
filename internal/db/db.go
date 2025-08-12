package db

import (
    "database/sql"
    _ "github.com/lib/pq"
    "log"
)

var DB *sql.DB

func Init() {
    var err error
    dsn := "host=localhost port=5432 user=youruser password=yourpass dbname=yourdb sslmode=disable"
    DB, err = sql.Open("postgres", dsn)
    if err != nil {
        log.Fatal("Failed to connect to DB:", err)
    }
    if err = DB.Ping(); err != nil {
        log.Fatal("Failed to ping DB:", err)
    }
}