package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	user     = "postgres"
	password = "Dilshod@2005"
	dbname   = "n9"
	port     = 5432
)

var db *sql.DB 

type Person struct {
	Id       int    `json:"id"`
	FullName string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func init() {
	
	dbInfo := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", host, user, password, dbname, port)
	var err error
	db, err = sql.Open("postgres", dbInfo)
	if err != nil {
		log.Fatal(err)
	}
}

func Register(w http.ResponseWriter, r *http.Request) {
	var person Person

	if err := json.NewDecoder(r.Body).Decode(&person); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var exists bool
	err := db.QueryRow("SELECT EXISTS (SELECT 1 FROM yengi WHERE fullname = $1 OR username = $2 OR email = $3)",
		person.FullName, person.Username, person.Email).Scan(&exists)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if exists {
		http.Error(w, "already registred", http.StatusConflict)
		return
	}

	_, err = db.Exec("INSERT INTO yengi (fullname, username, email, password) VALUES ($1, $2, $3, $4)",
		person.FullName, person.Username, person.Email, person.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Registered successfully"))
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	idUser := params.Get("id")
	id, err := strconv.Atoi(idUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var updatedPerson Person
	if err := json.NewDecoder(r.Body).Decode(&updatedPerson); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = db.Exec("UPDATE yengi SET fullname=$1, username=$2, email=$3, password=$4 WHERE id=$5",
		updatedPerson.FullName, updatedPerson.Username, updatedPerson.Email, updatedPerson.Password, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User updated successfully"))
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	idUser := params.Get("id")
	id, err := strconv.Atoi(idUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = db.Exec("DELETE FROM yengi WHERE id=$1", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User deleted successfully"))
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var newUser Person
	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err := db.Exec("INSERT INTO yengi (fullname, username, email, password) VALUES ($1, $2, $3, $4)",
		newUser.FullName, newUser.Username, newUser.Email, newUser.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User created successfully"))
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT id, fullname, username, email, password FROM yengi")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var users []Person
	for rows.Next() {
		var user Person
		if err := rows.Scan(&user.Id, &user.FullName, &user.Username, &user.Email, &user.Password); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		users = append(users, user)
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(users); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func main() {
	http.HandleFunc("/register", Register)
	http.HandleFunc("/update", UpdateUser)
	
	http.HandleFunc("/delete", DeleteUser)
	http.HandleFunc("/create", CreateUser)
	http.HandleFunc("/get", GetUsers)
	log.Fatal(http.ListenAndServe(":8080", nil))
}


