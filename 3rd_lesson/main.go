package main

import (
	"database/sql"
	"encoding/json"
	"log"

	"fmt"

	"net/http"

	_ "github.com/lib/pq"
)

type Book struct {
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

const (
	host     = "localhost"
	user     = "postgres"
	password = "Dilshod@2005"
	port     = 5432
	dbname   = "dilshod"
)

var db *sql.DB
var err error

func main() {
	dbInfo := fmt.Sprintf("host=%s user=%s password=%s port=%d dbname=%s sslmode=disable", host, user, password, port, dbname)
	db, err = sql.Open("postgres", dbInfo)
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/create", createData)
	http.HandleFunc("/change", modifyData)
	http.HandleFunc("/getdata", getData)
	http.HandleFunc("/delete", delete)

	fmt.Println("server runnning at :8080..")
	http.ListenAndServe(":8080", nil)

}

func jwtcheckAuthentification() {

}

func getData(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("select * from learn")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var books []Book
	for rows.Next() {
		var b Book
		if err := rows.Scan(&b.Id, &b.Title, &b.Author); err != nil {
			log.Fatal(err)
		}
		books = append(books, b)
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(books)

	if err != nil {
		log.Fatal(err)
	}
}

func createData(w http.ResponseWriter, r *http.Request) {
	var book = &Book{
		Title:  "atomic habit",
		Author: "George",
	}

	_, err := db.Exec("insert into learn(title, author) values($1, $2)", book.Title, book.Author)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(w, "Added the book brother...")

}

func modifyData(w http.ResponseWriter, r *http.Request) {
	var book Book
	book.Id = 1
	book.Title = "Dilshod"
	book.Author = "Dilshodbekkk"
	_, err = db.Exec("update learn set title=$1, author=$2 where id=$3", book.Title, book.Author, book.Id)
	fmt.Fprintf(w, "changed brother look at it...")
	if err != nil {
		log.Fatal(err)
	}

}

func delete(w http.ResponseWriter, r *http.Request) {
	var book Book
	book.Id = 2

	_, err = db.Exec("delete from learn where id=$1", book.Id)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(w, "bro deleted look...")
}
