package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		http.ServeFile(w, r, "registration.html")
	})

	http.HandleFunc("/dashboard", func(w http.ResponseWriter, r *http.Request) {

		email := r.FormValue("email")
		name := r.FormValue("name")

		if email == "" || name == "" {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}

		http.Redirect(w, r, "/dashboard.html?email="+email+"&name="+name, http.StatusSeeOther)
	})

	http.HandleFunc("/dashboard.html", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "dashboard.html")
	})

	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}
