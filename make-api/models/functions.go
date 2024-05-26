package postgres

import (
	"database/sql"
	"log"
	"net/http"
	models "student/models"
	

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type Student struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Age   int     `json:"age"`
	Grade int     `json:"grade"`
	Gpa   float32 `json:"gpa"`
}

func StudentGet(c *gin.Context) {
	db, err := postgres.Initialize()
	if err != nil {
		log.Fatal("smth wrong while initialize...", err)
	}
	rows, err := db.Query("SELECT * FROM student")
	if err != nil {
		log.Fatal("error while selecting infoo...", err)
	}
	var all = []Student{}
	for rows.Next() {
		var s Student
		if err = rows.Scan(&s.ID, &s.Name, &s.Age, &s.Grade, &s.Gpa); err != nil {
			log.Fatal("erorr while scanning from database...", err)
		}
		all = append(all, s)
	}
	c.IndentedJSON(http.StatusOK, all)
}

// func UpdateStudent(c *gin.Context){

// }

// func CreateStudent(c *gin.Context){

// }

// func DeleteStudent(c *gin.Context){

// }
