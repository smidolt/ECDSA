package internal

import (
    "database/sql"
    _ "github.com/lib/pq"
    "log"
)

func ConnectDB() (*sql.DB, error) {
    connStr := "postgres://newuser:newpassword@localhost/newdatabase?sslmode=disable"
    db, err := sql.Open("postgres", connStr)
    if err != nil {
        log.Fatalf("Failed to connect to the database: %v", err)
    }
    if err = db.Ping(); err != nil {
        log.Fatalf("Failed to ping the database: %v", err)
    }
    log.Println("Successfully connected to the database")
    return db, nil
}
