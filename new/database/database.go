package database

import (
	"database/sql"
	"fmt"
	"log"
	_"github.com/lib/pq"
	"github.com/gin-gonic/gin"
)

const (
	host     = "localhost"
	user     = "postgres"
	password = "Dilshod@2005"
	port     = 5432
	dbname   = "n9"
)

func Connection() *sql.DB{
	dbInfo := fmt.Sprintf("host=%s user=%s password=%s port=%d dbname=%s sslmode=disable", host, user, password, port, dbname)
	db, err := sql.Open("postgres", dbInfo)
	if err != nil{
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil{
		log.Fatal(err)
	}
	return db
	
}

func createTable(){
	con := Connection()
	
}