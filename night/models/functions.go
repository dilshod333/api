package models

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)
const (
	host="localhost"
	password="Dilshod@2005"
	user="postgres"
	dbname="learner"
	port=5432
)


func Connection()(*sql.DB, error){
	dbInfo := fmt.Sprintf("host=%s port=%d dbname=%s password=%s user=%s sslmode=disable", host, port, dbname,password, user)
	db,err := sql.Open("postgres", dbInfo)

	if err != nil{
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil{
		log.Fatal(err)
	}
	return db, err
}