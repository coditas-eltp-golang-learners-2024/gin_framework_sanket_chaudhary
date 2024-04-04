package config

import (
    "database/sql"
    "log"

    _ "github.com/go-sql-driver/mysql"
)

 //var DB *sql.DB

func ConnectionWithDb() (*sql.DB, error) {
    db, err := sql.Open("mysql", "root:S@nket123@tcp(localhost:3306)/mydb")
    if err != nil {
        log.Fatalf("Error opening database connection: %v", err)
        return nil, err
    }
    return db, nil
}
