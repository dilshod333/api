package handler

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)
// JWT KEY HALI ISHLAMIDI QOGANI ISHLIDI
var jwtKey = []byte("Dilshod")
type Quote struct {
	ID       int    `json:"id"`
	Text     string `json:"text"`
	Category string `json:"category"`
}

var quotes = []Quote{
	{ID: 1, Text: "you must be the change you wish to see in the world.", Category: "motivation"},
	{ID: 2, Text: "Spread love everywhere you go. ..", Category: "insparation"},
	{ID: 3, Text: "Darkness cannot drive out darkness: only light can do that. ...", Category: "motivation"},
	{ID: 4, Text: "Do one thing every day that scares you. - ..", Category: "insparation"},
}
type User struct{
	Id int 
	Name  string 
	Password int 

}

type Claims struct {
    Username string `json:"username"`
    jwt.StandardClaims
}

const (
	host     = "localhost"
	user     = "postgres"
	password = "Dilshod@2005"
	port     = 5432
	dbname   = "n9"
)

func Connect() *sql.DB {
	dbInfo := fmt.Sprintf("host=%s user=%s port=%d password=%s dbname=%s sslmode=disable", host, user, port, password, dbname)
	db, err := sql.Open("postgres", dbInfo)
	if err != nil {
		log.Fatal("error with connecting to the database...", err)
	}
	return db
}

func Register(c  *gin.Context){
	var user User


	if err := c.ShouldBindJSON(&user); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 
	}
	db := Connect()
	defer db.Close()
	_, err := db.Exec("insert into users(name, password) values($1, $2)", user.Name, user.Password)
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error registering user"})
		log.Fatal("error inserting user into table...", err)
		return 
	}


	claims := &Claims{  
		Username: user.Name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: jwt.TimeFunc().Add(time.Minute * 1).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims) // *jwt.Token qaytaradi..
	tokenString, err := token.SignedString(jwtKey)
	if err != nil{
		log.Fatal("error generating error..", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error generating jwt token"})
		return 
	}

	c.JSON(http.StatusOK, gin.H{"token":tokenString})

}
	

func AuthMiddleWare(c *gin.Context){  // *engine
	tokenString := c.GetHeader("Authorization")
	if tokenString == ""{
		c.JSON(http.StatusUnauthorized, gin.H{"error": "No authorization broo..."})
		c.Abort()
		return 
		
	}
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil{
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
        c.Abort()
        return
	}

	if !token.Valid{
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token broooo comeee onnn"})
        c.Abort()
        return
	}

	c.Set("user", claims.Username)
	c.Next()
}

func CreateTable() {
	db := Connect()
	fmt.Println("Connected successsfully to the database...")

	for _, quote := range quotes {
		fmt.Println(quote.Text, quote.Category)
		_, err := db.Exec("INSERT INTO justApi(text, category) VALUES($1, $2)", quote.Text, quote.Category)
		if err != nil {
			log.Fatal("error with inserting quote into the database:", err)
		}
	}

	fmt.Println("Quotes inserted successfully.")
}
func GetQuote(w *gin.Context) {
	db := Connect()

	rows, err := db.Query("select * from justApi")

	if err != nil {
		log.Fatal("Getting error....", err)
	}
	var getQuote []Quote

	for rows.Next() {
		var q Quote

		if err = rows.Scan(&q.ID, &q.Text, &q.Category); err != nil {
			log.Fatal("Scaninngg error...", err)
		}
		getQuote = append(getQuote, q)

	}
	w.JSON(http.StatusOK, getQuote)

}

func PostQuote(c *gin.Context) {
	db := Connect()
	var lastId int 
	err := db.QueryRow("select id from justApi order by id desc limit 1").Scan(&lastId)
	if err != nil && err != sql.ErrNoRows {
        log.Fatal("error retrieving last inserted ID:", err)
    }
	newId := lastId + 1
	
	
	newQuote := Quote{
		ID:       newId,
		Text:     "live like you die tommorow",
		Category: "Motivation",
	}

	quotes = append(quotes, newQuote)

	_, err = db.Exec("INSERT INTO justApi(id, text, category) VALUES($1, $2, $3)",newQuote.ID, newQuote.Text, newQuote.Category)
	if err != nil {
		log.Fatal("error inserting new quote into the database:", err)
	}

	c.IndentedJSON(http.StatusCreated, newQuote)
}

func AllCategoryQuote(w *gin.Context) {
	db := Connect()
	var categories []string
	rows, err := db.Query("SELECT  category FROM justApi")
	if err != nil {
		log.Fatal("error categories:", err)
	}
	defer rows.Close()

	for rows.Next() {
		var category string
		if err := rows.Scan(&category); err != nil {
			log.Fatal("error scanning category:", err)
		}
		categories = append(categories, category)
	}
	if err := rows.Err(); err != nil {
		log.Fatal("error iterating over category rows:", err)
	}

	w.JSON(http.StatusOK, categories)
}

func FilterQuote(c *gin.Context) {

	category := c.Param("category")
	var filterquote []Quote
	for _, q := range quotes {
		if q.Category == category {
			filterquote = append(filterquote, q)
		}
	}

	c.JSON(http.StatusOK, filterquote)

}

func DeleteQuote(c *gin.Context) {
	IdUser := c.Param("id")
	user, err := strconv.Atoi(IdUser)
	if err != nil {
		log.Fatal("error with converting string to integer...", err)
	}
	db := Connect()

	_, err = db.Exec("delete from justApi where id=$1", user)
	if err != nil {
		log.Fatal("error deleting quote from database:", err)
	}

	index := -1
	for idx, i := range quotes {
		if i.ID == user {
			index = idx
			break
		}
	}

	if index == -1 {

		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Quote not found"})
		return
	}
	quotes = append(quotes[:index], quotes[index+1:]...)

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Quote deleted successfully"})
}
