package main

import (
	"conn/handler"
	
	_ "github.com/lib/pq"
	"github.com/gin-gonic/gin"
)




func main() {
	
	router := gin.Default()
	router.GET("/register", handler.Register)
	router.GET("/quotes", handler.GetQuote)
	router.POST("/create", handler.PostQuote)
	router.GET("/category",handler.AuthMiddleWare, handler.AllCategoryQuote)
	router.GET("/quote/:category", handler.FilterQuote)
	router.DELETE("/delete/:id",handler.AuthMiddleWare, handler.DeleteQuote)
	router.Run(":80")
}
