package main

import (
	"conn/moduls"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)


// func main() {
	
// 	// // Insert a user
// 	// userID, err := moduls.InsertUser("John Doe", "john@example.com")
// 	// if err != nil {
// 	// 	log.Fatalf("Error inserting user: %v", err)
// 	// }
// 	// fmt.Printf("Inserted user with ID: %d\n", userID)

// 	// // Insert an order for the user
// 	// orderID, err := moduls.InsertOrder(userID, "Laptop", 1)
// 	// if err != nil {
// 	// 	log.Fatalf("Error inserting order: %v", err)
// 	// }
// 	// fmt.Printf("Inserted order with ID: %d\n", orderID)


// 	userOrders, err := moduls.GetUserOrders()
//     if err != nil {
//         log.Fatalf("Error getting user orders: %v", err)
//     }

//     // Print the retrieved data
//     fmt.Println("User Orders:")
//     for _, order := range userOrders {
//         fmt.Printf("User: %s, Product: %s, Amount: %d\n", order.UserName, order.Product, order.Amount)
//     }

// }



func main() {
    // Call GetUserOrders function
    userOrders, err := moduls.GetUserOrders()
    if err != nil {
        log.Fatalf("Error getting user orders: %v", err)
    }

    // Print the user orders
    fmt.Println("User Orders:")
    currentUsername := ""
    for _, order := range userOrders {
        // If the username has changed, print it
        if order.UserName != currentUsername {
            fmt.Printf("User: %s\n", order.UserName)
            currentUsername = order.UserName
        }
        // Print the order details
        fmt.Printf("  Product: %s, Amount: %d\n", order.Product, order.Amount)
    }
}
