// Step 4: Import necessary packages
package main

import (
    "encoding/json"
    "net/http"
    "time"

    "github.com/gorilla/mux"
    "github.com/dgrijalva/jwt-go"
)

// Step 5: Define a secret key for JWT
var secretKey = []byte("your_secret_key")

// Step 6: Define a User struct
type User struct {
    ID       int    `json:"id"`
    Username string `json:"username"`
    Password string `json:"password"`
}

// Step 7: Create a mock database of users
var users = []User{
    {ID: 1, Username: "user1", Password: "password1"},
    {ID: 2, Username: "user2", Password: "password2"},
}

// Step 8: Implement a route for user authentication and token generation
func loginHandler(w http.ResponseWriter, r *http.Request) {
    var user User
    json.NewDecoder(r.Body).Decode(&user)
    for _, u := range users {
        if u.Username == user.Username && u.Password == user.Password {
            // Step 9: Generate a JWT token
            token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
                "id":       u.ID,
                "username": u.Username,
                "exp":      time.Now().Add(time.Hour).Unix(),
            })
            tokenString, err := token.SignedString(secretKey)
            if err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
                return
            }
            json.NewEncoder(w).Encode(map[string]string{"token": tokenString})
            return
        }
    }
    http.Error(w, "Invalid username or password", http.StatusUnauthorized)
}

// Step 10: Implement a protected route
func profileHandler(w http.ResponseWriter, r *http.Request) {
    // Step 11: Extract JWT token from the request header
    tokenString := r.Header.Get("Authorization")
    if tokenString == "" {
        http.Error(w, "Authorization token not found", http.StatusUnauthorized)
        return
    }
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        return secretKey, nil
    })
    if err != nil {
        http.Error(w, err.Error(), http.StatusUnauthorized)
        return
    }
    if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
        // Step 12: Access user information from the JWT claims
        userID := int(claims["id"].(float64))
        username := claims["username"].(string)
        json.NewEncoder(w).Encode(map[string]interface{}{
            "message":  "Welcome to your profile",
            "userID":   userID,
            "username": username,
        })
    } else {
        http.Error(w, "Invalid token", http.StatusUnauthorized)
    }
}

func main() {
    // Step 13: Create an instance of the Gorilla mux router
    router := mux.NewRouter()

    // Step 14: Define route handlers
    router.HandleFunc("/login", loginHandler).Methods("POST")
    router.HandleFunc("/profile", profileHandler).Methods("GET")

    // Step 15: Start the server
    http.ListenAndServe(":3000", router)
}
