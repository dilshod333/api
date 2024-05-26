package main

import (
	"conn/models"
	"log"
	"fmt"
)

func main(){
	_, err := models.Connection()

	if err !=  nil{
		log.Fatal(err)
	}

	fmt.Println("Successfully connectedd...")
}