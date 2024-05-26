package main 
import (
	"github.com/lib/pq"
)

type Book struct{
	Title string `json:"title"`
	Description string `json:"description"`
	Price float32 `json:"price"`
	Quantity int `json:"quantity"`

}



func main(){

}