package main

import (
	"conn/models"

	// "github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)


func main() {
	r := models.SetupRouter()
	r.Run(":8080")
}
