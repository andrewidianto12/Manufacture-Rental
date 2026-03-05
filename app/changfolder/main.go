package main

import (
    "fmt"
    "log"
    "os"

    "github.com/andrewidianto12/miniproject-andre/database"
    "github.com/joho/godotenv"
)

func main() {
    loadEnv()

    db := database.InitDatabase()
    defer db.Close()

    addr := os.Getenv("APP_HOST") + ":" + os.Getenv("APP_PORT")
    fmt.Println("Database connection test successful")
    log.Default().Println("Application initialized on " + addr)
}

func loadEnv() {
    err := godotenv.Load("../../.env")
    if err != nil {
        log.Println("Warning: .env file not found, menggunakan sistem env")
    }
}
