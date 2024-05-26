package main

import (
	"log"
	"net/http"
    "conn/moduls"
  
)


func main(){
    http.HandleFunc("/login", moduls.Login)
    http.HandleFunc("/home", moduls.Home)
    http.HandleFunc("/refresh", moduls.Refresh)

    log.Fatal(http.ListenAndServe(":8080", nil))
}

