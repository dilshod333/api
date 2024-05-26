package postgres

import (
    "database/sql"
    "fmt"
    "log"
    "os"
    _ "github.com/lib/pq"
    "github.com/joho/godotenv"
)


func InitDB() (*sql.DB, error) {
    
    if err := godotenv.Load(); err != nil {
        log.Printf("No .env file found")
    }

   
    dbHost := os.Getenv("DB_HOST")
    dbPort := os.Getenv("DB_PORT")
    dbUser := os.Getenv("DB_USER")
    dbPass := os.Getenv("DB_PASS")
    dbName := os.Getenv("DB_NAME")

    
    connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
        dbHost, dbPort, dbUser, dbPass, dbName)

   
    db, err := sql.Open("postgres", connStr)
    if err != nil {
        return nil, fmt.Errorf("error opening database: %v", err)
    }

   
    if err := db.Ping(); err != nil {
        return nil, fmt.Errorf("error connecting to the database: %v", err)
    }

    return db, nil
}


