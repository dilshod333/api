package moduls

import (
	"conn/moduls/postgres"
	"fmt"
	"log"
)

func InsertUser(name, email string) (int, error) {
	db,err := postgres.InitDB()
	if err != nil{
		log.Fatal(err)
	}
	var userID int
	err = db.QueryRow("INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id", name, email).Scan(&userID)
	if err != nil {
		return 0, fmt.Errorf("error inserting user: %v", err)
	}
	return userID, nil
}


func InsertOrder(userID int, product string, amount int) (int, error) {
	db,err := postgres.InitDB()
	if err != nil{
		log.Fatal(err)
	}
	var orderID int
	err = db.QueryRow("INSERT INTO orders (user_id, product, amount) VALUES ($1, $2, $3) RETURNING id", userID, product, amount).Scan(&orderID)
	if err != nil {
		return 0, fmt.Errorf("error inserting order: %v", err)
	}
	return orderID, nil
}


func GetUserOrders() ([]UserOrder, error) {
    db, err := postgres.InitDB()
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    rows, err := db.Query("SELECT u.name AS user_name, o.product, o.amount FROM users u JOIN orders o ON u.id = o.user_id")
    if err != nil {
        return nil, fmt.Errorf("error querying user orders: %v", err)
    }
    defer rows.Close()

    var userOrders []UserOrder
    for rows.Next() {
        var userOrder UserOrder
        if err := rows.Scan(&userOrder.UserName, &userOrder.Product, &userOrder.Amount); err != nil {
            return nil, fmt.Errorf("error scanning row: %v", err)
        }
        userOrders = append(userOrders, userOrder)
    }
    if err := rows.Err(); err != nil {
        return nil, fmt.Errorf("error iterating over rows: %v", err)
    }

    return userOrders, nil
}


type UserOrder struct {
    UserName string
    Product  string
    Amount   int
}
