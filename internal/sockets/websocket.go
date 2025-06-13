package sockets

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan interface{})

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func HandleConnections(c *gin.Context) {
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Printf("WebSocket error: %v", err)
		return
	}
	defer ws.Close()

	clients[ws] = true

	for {
		var msg interface{}
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("WebSocket closed: %v", err)
			delete(clients, ws)
			break
		}
	}
}

func HandleMessages() {
	for {
		msg := <-broadcast
		for client := range clients {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Printf("WebSocket send error: %v", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}

func SendMessage(message interface{}) {
	broadcast <- message
}
