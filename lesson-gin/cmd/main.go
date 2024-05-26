// package main

// import (
//     "database/sql"
//     "fmt"
//     "log"

//     _ "github.com/lib/pq"
// )


// type User struct {
//     ID   int
//     Name string
// }


// type Post struct {
//     ID      int
//     Title   string
//     Content string
//     UserID  int // Foreign key
// }

// func main() {
   
//     db, err := sql.Open("postgres", "postgres://postgres:Dilshod@2005@localhost/n9?sslmode=disable")
//     if err != nil {
//         log.Fatal("Failed to connect database:", err)
//     }
//     defer db.Close()

  
//     _, err = db.Exec("INSERT INTO users (name) VALUES ($1), ($2), ($3)", "Alice", "Bob", "Charlie")
//     if err != nil {
//         log.Fatal("Failed to insert users:", err)
//     }

//     _, err = db.Exec("INSERT INTO posts (title, content, user_id) VALUES ($1, $2, $3), ($4, $5, $6)", "Post 1", "Content 1", 1, "Post 2", "Content 2", 2)
//     if err != nil {
//         log.Fatal("Failed to insert posts:", err)
//     }

  
//     rows, err := db.Query("SELECT p.title, p.content, u.name FROM posts p JOIN users u ON p.user_id = u.id")
//     if err != nil {
//         log.Fatal("Failed to query posts:", err)
//     }
//     defer rows.Close()

//     // Iterate over the rows and print the results
//     for rows.Next() {
//         var title, content, userName string
//         if err := rows.Scan(&title, &content, &userName); err != nil {
//             log.Fatal("Failed to scan row:", err)
//         }
//         fmt.Printf("Post Title: %s\nContent: %s\nUser Name: %s\n\n", title, content, userName)
//     }
// }


package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func main() {
	db, err := sql.Open("postgres", "postgres://username:password@localhost/crud_api?sslmode=disable")
	if err != nil {
		log.Fatal("Failed to connect database:", err)
	}
	defer db.Close()

	router := gin.Default()

	router.GET("/users", func(c *gin.Context) {
		var users []User
		rows, err := db.Query("SELECT id, username, password FROM users")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query users"})
			return
		}
		defer rows.Close()

		for rows.Next() {
			var user User
			if err := rows.Scan(&user.ID, &user.Username, &user.Password); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan users"})
				return
			}
			users = append(users, user)
		}

		c.JSON(http.StatusOK, users)
	})

	router.POST("/users", func(c *gin.Context) {
		var user User
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON payload"})
			return
		}

		user.ID++

		_, err := db.Exec("INSERT INTO users (username, password) VALUES ($1, $2)", user.Username, user.Password)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert user"})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"message": "User created", "user": user})
	})

	router.PUT("/users/:id", func(c *gin.Context) {
		id := c.Param("id")
		var user User
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON payload"})
			return
		}

		userID, _ := strconv.Atoi(id)
		_, err := db.Exec("UPDATE users SET username=$1, password=$2 WHERE id=$3", user.Username, user.Password, userID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "User updated", "user": user})
	})

	router.DELETE("/users/:id", func(c *gin.Context) {
		id := c.Param("id")

		_, err := db.Exec("DELETE FROM users WHERE id=$1", id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "User deleted", "id": id})
	})

	router.POST("/login", func(c *gin.Context) {
		var loginReq struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}
		if err := c.BindJSON(&loginReq); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON payload"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
	})

	fmt.Println("Server running on port 8080...")
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
