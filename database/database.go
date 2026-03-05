package database

import (
    "database/sql"
    "fmt"
    "os"
    "time"

    _ "github.com/jackc/pgx/v5/stdlib"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

func InitDatabase() *sql.DB {
    dsn := os.Getenv("DB_DSN")

    if dsn == "" {
        panic("DB_DSN is required")
    }

    db, err := sql.Open("pgx", dsn)
    if err != nil {
        fmt.Println("error InitDatabase", err)
        panic(err)
    }
    db.SetConnMaxLifetime(time.Minute * 3)
    db.SetMaxOpenConns(10)
    db.SetMaxIdleConns(10)

    err = db.Ping()
    if err != nil {
        fmt.Println("Failed to connect to the database")
        panic(err)
    }

    fmt.Println("Connect to the database succesfully (postgres)")

    return db
}

func InitDatabaseWithGORM() *gorm.DB {
    dsn := os.Getenv("DB_DSN")
    if dsn == "" {
        panic("DB_DSN is required")
    }
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        fmt.Println("Failed to connect to the database")
        panic(err)
    }

    if os.Getenv("DB_DEBUG") == "true" {
        return db.Debug()
    }

    return db
}
