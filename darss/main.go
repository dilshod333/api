package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const (
	name     = "Dilshod"
	password = "Dilshod@2005"
)

func main() {
	http.HandleFunc("/login", Login)

	fmt.Println("Server is listening on port 8080...")

	http.ListenAndServe(":8080", nil)
}

type Person struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	var person Person
	check := json.NewDecoder(r.Body).Decode(&person)
	if check != nil {
		log.Fatal(check)

	}
	if person.Name == name && person.Password == password {
		fmt.Fprintf(w, "righttt")
	}else{
		fmt.Fprintf(w, "wrong")

	}
	// fmt.Printf("Name: %s, Password: %s\n", person.Name, person.Password)

}
