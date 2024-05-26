package postgres

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	_"github.com/lib/pq"
	"github.com/joho/godotenv"
	// _ "github.com/lib/pq"
)


func Connection() (*sql.DB, error) {
	if err := godotenv.Load(); err != nil{
		return nil,err 
	}
	dbHost := os.Getenv("DBHOST")
	dbUser := os.Getenv("DBUSER")
	dbPassword := os.Getenv("DBPASSWORD")
	dbPort := os.Getenv("DBPORT")
	dbName := os.Getenv("DBNAME")

	dbInfo := fmt.Sprintf("host=%s user=%s password=%s port=%s dbname=%s",dbHost, dbUser, dbPassword, dbPort, dbName)
	db, err := sql.Open("postgres", dbInfo)
	if err != nil{
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil{
		log.Fatal(err)
	}
	return db, nil

}
