package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)
type Student struct{
	Id int `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

type Course struct{
	Id int `json:"id"`
	Title string `json:"title"`
	Price float32 `json:"price"`
	StudentCount int `json:"student_count"`
	Created_at time.Time `json:"created_at"`
}
var db *sql.DB
const (
	host = "localhost"
	port = 5432
	user = "postgres"
	password = "Dilshod@2005"
	dbname = "n9"

)

var data = [][]string{
	{"John", "john@example.com", "2000-01-01"},
	{"Alice", "alice@example.com", "2001-02-02"},
	{"Bob", "bob@example.com", "2002-03-03"},
	{"Carol", "carol@example.com", "2003-04-04"},
	{"David", "david@example.com", "2004-05-05"},
}

var course = [][]interface{}{
    {"Math", 100.23, 11, "2006-10-14"},
    {"English", 120.50, 15, "2007-11-15"},
    {"Science", 90.75, 10, "2008-12-16"},
    {"History", 80.30, 8, "2009-01-17"},
    {"Geography", 110.80, 12, "2010-02-18"},
}
var err error
func main(){
	dbInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err  = sql.Open("postgres", dbInfo)
	if err != nil{
		log.Fatal("erorr to connecting database....", err)
	}
	
	router := gin.Default()
	router.GET("/get", getStudent)
	router.GET("/get/course", getCourse)
	// router.PUT("/change", changeStudent)
	router.Run(":8080")
	
}

func getStudent(c *gin.Context) {
	
	rows, err := db.Query("SELECT * FROM student")
	if err != nil {
		log.Fatal(err)
	}

	var students []Student

	
	for rows.Next() {
		var st Student
	
		if err := rows.Scan(&st.Id, &st.Name, &st.Email, &st.CreatedAt, &st.UpdatedAt, &st.DeletedAt); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "failed to scan row"})
			return
		}
		
		students = append(students, st)
	}

	
	if err := rows.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "database iteration error"})
		return
	}

	
	rows.Close()

	
	c.JSON(http.StatusOK, students)
}
	


func InsertData(){


	for _, d := range data {
		_, err := db.Exec(`INSERT INTO Student (Name, Email, CreatedAt, UpdatedAt, DeletedAt) 
						   VALUES ($1, $2, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, $3)`,
			d[0], d[1], d[2])
		if err != nil {
			log.Fatalf("Failed to insert data for %s: %v", d[0], err)
		}
	}

	fmt.Println("Insertion successful")
}

func InsertCourse() {
    for _, c := range course {
     
        _, err := db.Exec(`INSERT INTO Course (Title, Price, StudentCount, Createdat) VALUES ($1, $2, $3, $4)`, c[0], c[1], c[2], c[3])

        if err != nil {
            log.Fatalf("Failed to insert course: %v", err)
        }
    }

    fmt.Println("Course insertion successful")
}


func getCourse(c *gin.Context){
	rows, err := db.Query("select * from course")
	if err != nil{
		log.Fatal("error while getting course...",err)

	}

	for rows.Next(){
		var c Course
		if err = rows.Scan(&)
	}
}