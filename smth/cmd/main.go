package main

import (
	// "conn/modules/postgres"
	// "conn/modules/postgres"
	"fmt"
	"just/modules/postgres"
	"log"
	// "just/postgres"
)

func main() {
	db, err := postgres.Connection()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("It is worked...", db)
}
