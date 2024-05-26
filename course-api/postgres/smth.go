package postgres

import "fmt"

type People struct {
	Name  string
	Age   int
	Email string
}

var people = []People{
	{Name: "Dilshod", Age: 19, Email: "Dilshod@2005"},
}

func Work() {
	var s = people
	fmt.Println(s)
}
