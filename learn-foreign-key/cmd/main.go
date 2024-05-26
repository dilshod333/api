package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)
const (
	host = "localhost"
	port = 5432
	user = "postgres"
	password = "Dilshod@2005"
	dbname = "learner"
)
func main() {
	// Connect to the PostgreSQL database
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}
	defer db.Close()

	// Create schools table
	if _, err := db.Exec(`CREATE TABLE IF NOT EXISTS schools (
		id SERIAL PRIMARY KEY,
		name TEXT NOT NULL
	)`); err != nil {
		log.Fatal("Error creating schools table:", err)
	}

	// Create students table with a foreign key constraint
	if _, err := db.Exec(`CREATE TABLE IF NOT EXISTS students (
		id SERIAL PRIMARY KEY,
		name TEXT NOT NULL,
		school_id INT NOT NULL,
		FOREIGN KEY (school_id) REFERENCES schools(id)
	)`); err != nil {
		log.Fatal("Error creating students table:", err)
	}

	log.Println("Tables created successfully")

	// Insert data into the schools table
	schoolID1, err := insertSchool(db, "Oak Elementary School")
	if err != nil {
		log.Fatal("Error inserting school:", err)
	}

	schoolID2, err := insertSchool(db, "Pine Middle School")
	if err != nil {
		log.Fatal("Error inserting school:", err)
	}

	// Insert data into the students table
	_, err = insertStudent(db, "Alice", schoolID1)
	if err != nil {
		log.Fatal("Error inserting student:", err)
	}

	_, err = insertStudent(db, "Bob", schoolID1)
	if err != nil {
		log.Fatal("Error inserting student:", err)
	}

	_, err = insertStudent(db, "Charlie", schoolID2)
	if err != nil {
		log.Fatal("Error inserting student:", err)
	}

	log.Println("Data inserted successfully")
}

func insertSchool(db *sql.DB, name string) (int, error) {
	var id int
	err := db.QueryRow("INSERT INTO schools(name) VALUES($1) RETURNING id", name).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func insertStudent(db *sql.DB, name string, schoolID int) (int, error) {
	var id int
	err := db.QueryRow("INSERT INTO students(name, school_id) VALUES($1, $2) RETURNING id", name, schoolID).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}
