package postgres

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	dbname   = "very"
	password = "Dilshod@2005"
)

func Connection() (*sql.DB, error) {
	dbInfo := fmt.Sprintf("host=%s user=%s port=%d dbname=%s password=%s sslmode=disable", host, user, port, dbname, password)
	db, err := sql.Open("postgres", dbInfo)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	return db, err
}

func CreateTable() {
	db, err := Connection()
	if err != nil {
		log.Fatal(err)
	}

	studentTableQuery := `
		CREATE TABLE IF NOT EXISTS student(
			id SERIAL PRIMARY KEY,
			student_id INT UNIQUE NOT NULL,
			name TEXT,
			age INT,
			email TEXT
		);
	`
	_, err = db.Exec(studentTableQuery)
	if err != nil {
		log.Fatal(err)
	}

	courseTableQuery := `
		CREATE TABLE IF NOT EXISTS course(
			course_id SERIAL PRIMARY KEY,
			course_name TEXT
		);
	`
	_, err = db.Exec(courseTableQuery)
	if err != nil {
		log.Fatal(err)
	}

	genreTableQuery := `
		CREATE TABLE IF NOT EXISTS genre(
			genre_id SERIAL PRIMARY KEY,
			student_id INT,
			course_id INT,
			FOREIGN KEY (student_id) REFERENCES student(student_id),
			FOREIGN KEY (course_id) REFERENCES course(course_id)
		);
	`
	_, err = db.Exec(genreTableQuery)
	if err != nil {
		log.Fatal(err)
	}
}

func InsertStudent(studentID int, name string, age int, email string) {
	db, err := Connection()
	if err != nil {
		log.Fatal(err)
	}

	insertQuery := `
		INSERT INTO student (student_id, name, age, email)
		VALUES ($1, $2, $3, $4)
	`
	_, err = db.Exec(insertQuery, studentID, name, age, email)
	if err != nil {
		log.Fatal(err)
	}
}

func InsertCourse(courseName string) int {
	db, err := Connection()
	if err != nil {
		log.Fatal(err)
	}

	var courseID int
	insertQuery := `
		INSERT INTO course (course_name)
		VALUES ($1)
		RETURNING course_id
	`
	err = db.QueryRow(insertQuery, courseName).Scan(&courseID)
	if err != nil {
		log.Fatal(err)
	}
	return courseID
}

func InsertGenre(studentID int, courseID int) {
	db, err := Connection()
	if err != nil {
		log.Fatal(err)
	}

	insertQuery := `
		INSERT INTO genre (student_id, course_id)
		VALUES ($1, $2)
	`
	_, err = db.Exec(insertQuery, studentID, courseID)
	if err != nil {
		log.Fatal(err)
	}
}



func FetchStudentCourses() {
	db, err := Connection()
	if err != nil {
		log.Fatal(err)
	}

	query := `
		SELECT s.name, s.age, s.email, c.course_name
		FROM student s
		JOIN genre g ON s.student_id = g.student_id
		JOIN course c ON g.course_id = c.course_id;
	`

	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var name string
		var age int
		var email, courseName string

		err := rows.Scan(&name, &age, &email, &courseName)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Name: %s, Age: %d, Email: %s, Course: %s\n", name, age, email, courseName)
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
}
