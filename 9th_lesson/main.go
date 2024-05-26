package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Quote struct {
	ID       int    `json:"id"`
	Text     string `json:"text"`
	Category string `json:"category"`
}

var quotes = []Quote{
	{ID: 1, Text: "The greatest glory in living lies not in never falling, but in rising every time we fall.", Category: "Inspirational"},
	{ID: 2, Text: "The way to get started is to quit talking and begin doing.", Category: "Motivational"},
	{ID: 3, Text: "Life is what happens when you're busy making other plans.", Category: "Life"},
	{ID: 4, Text: "The only impossible journey is the one you never begin.", Category: "Inspirational"},
}


var count = len(quotes)
func main() {
	r := gin.Default()
	r.GET("/quote", quoteShow)
	r.GET("/quote/:category", Filter)
	r.GET("/category", Category)
	r.POST("create", Create)
	r.Run(":80")
}

func quoteShow(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, quotes)
}

func Category(c *gin.Context) {
	lst := []string{}
	for _, i := range quotes {
		lst = append(lst, i.Category)
	}
	c.JSON(http.StatusOK, lst)
}

func Filter(c *gin.Context) {
	var filtedQuotes []Quote
	getIt := c.Param("category")
	for _, i := range quotes {
		if getIt == i.Category {
			filtedQuotes = append(filtedQuotes, i)
		}
	}
	if len(filtedQuotes) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"category": "404 Not Found"})
	} else {
		c.JSON(http.StatusOK, filtedQuotes)
	}
}


func Create(c *gin.Context){
	count++
	var newquote = []Quote{
		{ID: count, Text: "live like you are gonna die tommorow", Category: "motivation"},
	}
	quotes = append(quotes, newquote...)
	c.JSON(http.StatusOK, newquote)
}