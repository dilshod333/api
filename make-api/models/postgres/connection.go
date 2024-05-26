package postgres

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	user     = "postgres"
	password = "Dilshod@2005"
	port     = 5432
	dbname   = "learner"
)

type Student struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Age   int     `json:"age"`
	Grade int     `json:"grade"`
	Gpa   float32 `json:"gpa"`
}

var student = []Student{
	{ID: 1, Name: "Alibek", Age: 19, Grade: 1, Gpa: 4.5},
	{ID: 2, Name: "Dilshodbek", Age: 19, Grade: 1, Gpa: 4.5},
	{ID: 3, Name: "Salimbek", Age: 22, Grade: 4, Gpa: 4.0},
	{ID: 4, Name: "Ibrohim", Age: 21, Grade: 3, Gpa: 4.5},
	{ID: 5, Name: "Sardor", Age: 23, Grade: 3, Gpa: 3.0},
}

func Initialize() (*sql.DB, error) {
	dbInfo := fmt.Sprintf("host=%s user=%s password=%s port=%d dbname=%s sslmode=disable", host, user, password, port, dbname)
	db, err := sql.Open("postgres", dbInfo)
	if err != nil {
		log.Fatal("error while connecting....", err)
	}
	return db, nil
}

func InsertData() {
	db, err := Initialize()

	if err != nil {
		log.Fatal(err)
	}

	stmt, err := db.Prepare("insert into student(id, name, age, grade, gpa) values($1, $2, $3, $4, $5)")
	if err != nil {
		log.Fatal("error while preparing...", err)
	}
	defer stmt.Close()
	for _, m := range student {
		_, err := stmt.Exec(m.ID, m.Name, m.Age, m.Grade, m.Gpa)
		if err != nil {
			log.Fatal("error while inserting data...", err)
		}
	}

	fmt.Println("successfully stored data....")
}
