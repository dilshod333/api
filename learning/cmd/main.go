// package main 

import (
	"fmt"
	"log"
	"new/moduls/postgres"
	"new/moduls"
)

func main() {
	res, err := postgres.Initialize()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("it is connected...",res)

}
