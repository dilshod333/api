package models

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Student struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func GetAllInfo(c *gin.Context) {
	db, err := Connection()
	if err != nil {
		log.Fatal(err)
	}
	rows, err := db.Query("select * from  person")
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {

	}
}

func UpdateInfo(c *gin.Context) {
	userId := c.Param("id")
	id, err := strconv.Atoi(userId)
	if err != nil{
		log.Fatal(err)
	}
	
	db, err := Connection()
	if err != nil{
		log.Fatal(err)
	}
	defer db.Close()
	rows, err := db.Query("select * from somewhere")
}

func Login(c *gin.Context) {

}

func Register(c *gin.Context) {

}

func DeleteStudent(c *gin.Context) {

}

func CreateStudent(c *gin.Context) {

}
