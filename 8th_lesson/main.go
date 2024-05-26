package main

import (
	"log"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var (
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	// Store active connections and chat rooms
	connections = make(map[*websocket.Conn]bool)
	rooms       = make(map[string][]*websocket.Conn)
	mutex       = sync.Mutex{}
)

func main() {
	r := gin.Default()

	// Register routes
	r.POST("/register", registerHandler)
	r.GET("/ws/:roomID/:username", websocketHandler)

	// Start server
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}

// Handler to register a new user (not implemented in this example)
func registerHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}

// WebSocket handler
func websocketHandler(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("Failed to upgrade to WebSocket:", err)
		return
	}
	defer conn.Close()

	// Extract room ID and username from URL parameters
	roomID := c.Param("roomID")
	username := c.Param("username")
	_ = username

	// Register connection
	mutex.Lock()
	connections[conn] = true
	rooms[roomID] = append(rooms[roomID], conn)
	mutex.Unlock()

	// Listen for incoming messages
	for {
		var msg map[string]string
		if err := conn.ReadJSON(&msg); err != nil {
			log.Println("Failed to read message:", err)
			break
		}

		// Broadcast message to all connections in the room
		mutex.Lock()
		for _, c := range rooms[roomID] {
			if err := c.WriteJSON(msg); err != nil {
				log.Println("Failed to send message:", err)
				continue
			}
		}
		mutex.Unlock()
	}

	// Remove connection when the loop exits
	mutex.Lock()
	delete(connections, conn)
	rooms[roomID] = removeConnection(rooms[roomID], conn)
	mutex.Unlock()
}

// Helper function to remove a connection from a slice
func removeConnection(slice []*websocket.Conn, conn *websocket.Conn) []*websocket.Conn {
	for i, c := range slice {
		if c == conn {
			return append(slice[:i], slice[i+1:]...)
		}
	}
	return slice
}
