package main 

import(
	"fmt"
)
type User struct{
	ID int 
	Name string 
	Email string 
	Password string
	Info []string
	Address *Address 

}

type Address struct{
	location string 
	postal_code int 
}

func main(){
	var user = User{
		ID: 1,
		Name: "Dilshod",
		Email: "Dilshod@2005",
		Password: "Salom@20000",
		Info: []string{
			"software enginner",
			"nothing",
		},
		Address: &Address{
			location: "Navoi",
			postal_code: 441,
		},
	}
	fmt.Print(user)
}