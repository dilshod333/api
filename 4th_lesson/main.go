package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"

	_ "github.com/lib/pq"
)

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

const (
	host      = "localhost"
	port      = 5432
	user      = "postgres"
	password  = "Dilshod@2005"
	dbname    = "n9"
	secretKey = "secret" // Change this to a secure random value in production
	tokenTTL  = 24 * time.Hour
)

var db *sql.DB

func main() {
	initDB()

	http.HandleFunc("/register", registerHandler)
	http.HandleFunc("/login", loginHandler)

	// Protected routes
	http.HandleFunc("/api/books", authMiddleware(getBooks))
	http.HandleFunc("/api/books/create", authMiddleware(createBook))
	http.HandleFunc("/api/books/update", authMiddleware(updateBook))
	http.HandleFunc("/api/books/delete", authMiddleware(deleteBook))

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func initDB() {
	dbInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	var err error
	db, err = sql.Open("postgres", dbInfo)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	createTable()
}

func createTable() {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS books (
			id SERIAL PRIMARY KEY,
			title VARCHAR(255) NOT NULL,
			author VARCHAR(255) NOT NULL
		);
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			username VARCHAR(255) UNIQUE NOT NULL,
			password VARCHAR(255) NOT NULL
		);
	`)
	if err != nil {
		log.Fatal(err)
	}
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Check if the username already exists
	var existingUser User
	err := db.QueryRow("SELECT id FROM users WHERE username = $1", user.Username).Scan(&existingUser.ID)
	switch {
	case err == sql.ErrNoRows:
		// Username does not exist, proceed with registration
	case err != nil:
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	default:
		http.Error(w, "Username already exists", http.StatusBadRequest)
		return
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Failed to hash password", http.StatusInternalServerError)
		return
	}

	// Insert the new user into the database
	_, err = db.Exec("INSERT INTO users(username, password) VALUES($1, $2)", user.Username, string(hashedPassword))
	if err != nil {
		http.Error(w, "Failed to register user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, "User registered successfully")
}


func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var user struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var storedPassword string
	var userID int

	row := db.QueryRow("SELECT id, password FROM users WHERE username = $1", user.Username)
	err := row.Scan(&userID, &storedPassword)
	if err != nil {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(user.Password))
	if err != nil {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	tokenString, err := generateToken(userID, user.Username)
	if err != nil {
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": tokenString})
}

func authMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			http.Error(w, "Authorization header missing", http.StatusUnauthorized)
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(secretKey), nil
		})
		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		if !token.Valid {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	}
}

func generateToken(userID int, username string) (string, error) {
	expirationTime := time.Now().Add(tokenTTL)
	claims := jwt.MapClaims{
		"user_id":  userID,
		"username": username,
		"exp":      expirationTime.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secretKey))
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	// Implement your getBooks logic here
	fmt.Fprintf(w, "salom dilshoddd......")
}

func createBook(w http.ResponseWriter, r *http.Request) {
	// Implement your createBook logic here
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	// Implement your updateBook logic here
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	// Implement your deleteBook logic here
}
