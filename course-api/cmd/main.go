package main

import (
	conn "conn/postgres"
	"log"
	"fmt"
)

func main(){
	_, err := conn.Connection()
	if err != nil{
		log.Fatal(err)
	}

	fmt.Println("successfully connected...")
	// conn.CreateTable()

	// conn.InsertStudent(3, "Dilshod", 19, "Dilshod@example.com")
	// conn.InsertStudent(4, "Dostonbek", 22, "Doston@example.com")

	// courseID1 := conn.InsertCourse("English")
	// courseID2 := conn.InsertCourse("Math")

	// conn.InsertGenre(3, courseID1)
	// conn.InsertGenre(4, courseID2)

	conn.FetchStudentCourses()
}