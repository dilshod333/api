package main

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Quote represents a single quote
type Quote struct {
	Text     string `json:"text"`
	Category string `json:"category"`
}

// List of quotes
var quotes = []Quote{
	{Text: "The greatest glory in living lies not in never falling, but in rising every time we fall.", Category: "Inspirational"},
	{Text: "The way to get started is to quit talking and begin doing.", Category: "Motivational"},
	{Text: "Life is what happens when you're busy making other plans.", Category: "Life"},
	{Text: "The only impossible journey is the one you never begin.", Category: "Inspirational"},
	// Add more quotes...
}

func main() {
	r := gin.Default()

	// Endpoint to retrieve a random quote
	r.GET("/quote", getRandomQuoteHandler)

	// Endpoint to retrieve a random quote from a specific category
	r.GET("/quote/:category", getRandomQuoteByCategoryHandler)

	// Endpoint to list available categories
	r.GET("/categories", listCategoriesHandler)

	// Start server
	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}

// Handler to retrieve a random quote
func getRandomQuoteHandler(c *gin.Context) {
	rand.Seed(time.Now().Unix())
	quote := quotes[rand.Intn(len(quotes))]
	c.JSON(http.StatusOK, quote)
}

// Handler to retrieve a random quote from a specific category
func getRandomQuoteByCategoryHandler(c *gin.Context) {
	category := c.Param("category")
	var filteredQuotes []Quote
	for _, q := range quotes {
		if q.Category == category {
			filteredQuotes = append(filteredQuotes, q)
		}
	}
	if len(filteredQuotes) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		return
	}
	rand.Seed(time.Now().Unix())
	quote := filteredQuotes[rand.Intn(len(filteredQuotes))]
	c.JSON(http.StatusOK, quote)
}

// Handler to list available categories
func listCategoriesHandler(c *gin.Context) {
	categories := make(map[string]bool)
	for _, q := range quotes {
		categories[q.Category] = true
	}
	var categoryList []string
	for category := range categories {
		categoryList = append(categoryList, category)
	}
	c.JSON(http.StatusOK, categoryList)
}
