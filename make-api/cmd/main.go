package main

import (
	
	"student/models"

	"github.com/gin-gonic/gin"
)

func main(){
	router := gin.Default()
	router.GET("/student", models.StudentGet)
	// router.POST("/create", models.UpdateStudent)
	// router.PUT("/update/:id", models.CreateStudent)
	// router.DELETE("/delete/:id", models.DeleteStudent)

	router.Run(":8080")
}
