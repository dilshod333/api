package postgres

import (
	"database/sql"
	"os"
	"github.com/joho/godotenv"
	// _ "github.com/lib/pq"
)

func Initialize() (*sql.DB, error) {

	if err := godenv.Load(); err != nil {
		return nil
	}

	dbHost := os.Getenv("DBHOST")
	dbUser := os.Getenv
	dbPort := os.Getenv("DBPORT")
	// dbName := os.Getenv("DBNAME")
	dbPassword := os.Getenv("DBPASSWORD")

	dbInfo := "host" dbHost + " port=" + dbPort + " user=" + dbUser + " dbname=" + dbName + " password=" + dbPassword + " sslmode=disable"

	db, err := sql.Open("postgres", dbInfo)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db
}
