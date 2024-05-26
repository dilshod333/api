package models

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"
	"github.com/gin-gonic/gin"
	"github.com/dgrijalva/jwt-go"
	_ "github.com/lib/pq"

)

type User struct {
	ID       uint8
	Name     string
	Password string
	Email    string
}

const (
	host     = "localhost"
	password = "Dilshod@2005"
	port     = 5432
	dbname   = "n9"
	user     = "postgres"
	secret   = "your_secret_key"
)

func Connection() *sql.DB {
	dbInfo := fmt.Sprintf("host=%s port=%d password=%s dbname=%s  user=%s sslmode=disable", host, port, password, dbname, user)
	db, err := sql.Open("postgres", dbInfo)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Error to connection to database....", err)
	}

	return db
}

func GetAllData() {
	db := Connection()

	rows, err := db.Query("select * from practise")
	if err != nil {
		log.Fatal(err)
	}
	var userInfo []User
	defer db.Close()
	for rows.Next() {
		var user User

		if err = rows.Scan(&user.ID, &user.Name, &user.Password, &user.Email); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("id=%d name=%s password=%s email=%s\n\n", user.ID, user.Name, user.Password, user.Email)
		userInfo = append(userInfo, user)
	}

	fmt.Println("it is userInfo...", userInfo)
}

func GenerateToken(user User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = user.ID
	claims["name"] = user.Name
	claims["email"] = user.Email
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // Token expires in 24 hours

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// Login handles the user login functionality
func Login(c *gin.Context) {
	var user User
	// Bind JSON request body to User struct
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Query the database to fetch the user's password
	db := Connection()
	defer db.Close()
	row := db.QueryRow("SELECT password FROM practise WHERE name = $1", user.Name)

	var passwordFromDB string
	// Check if the user exists
	err := row.Scan(&passwordFromDB)
	if err != nil {
		// User not found or error occurred
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Compare the passwords
	if user.Password != passwordFromDB {
		// Passwords don't match
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Generate JWT token for the authenticated user
	token, err := GenerateToken(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	// Return the generated token as a response
	c.JSON(http.StatusOK, gin.H{"token": token})
}


func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			c.Abort()
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Check the token signing method
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(secret), nil
		})

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			// You can access the claims like this:
			// userID := claims["id"].(uint8)
			// userName := claims["name"].(string)
			// userEmail := claims["email"].(string)
			c.Set("userID", claims["id"])
			c.Next()
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}
	}
}

func ProtectedRoute(c *gin.Context) {
	// This handler will only be called if the token is valid
	userID, _ := c.Get("userID")
	c.JSON(http.StatusOK, gin.H{"message": "You are authorized", "userID": userID})
}

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/login", Login)

	// Use the AuthMiddleware for routes that require authentication
	r.Use(AuthMiddleware())
	{
		r.GET("/protected", ProtectedRoute)
	}

	return r
}

