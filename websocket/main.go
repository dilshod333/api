// main.go
package main

import (
    "github.com/gin-gonic/gin"
    "github.com/gorilla/websocket"
    "net/http"
)

var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
}

func main() {
    router := gin.Default()

    router.GET("/", func(c *gin.Context) {
        http.ServeFile(c.Writer, c.Request, "index.html")
    })

    router.GET("/ws", handleWebSocket)

    router.Run(":8080")
}

func handleWebSocket(c *gin.Context) {
    conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    defer conn.Close()

    for {
        _, msg, err := conn.ReadMessage()
        if err != nil {
            break
        }
      
        broadcastMessage(msg)
    }
}

var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan []byte)

func broadcastMessage(message []byte) {
    for client := range clients {
        err := client.WriteMessage(websocket.TextMessage, message)
        if err != nil {
            client.Close()
            delete(clients, client)
        }
    }
}

