package websocket

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var (
	Upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true // Allow all origins
		},
	}
	GlobalWebSocketCon *websocket.Conn
)

func SendWebSocketMessage(message []byte) {
	if GlobalWebSocketCon != nil {
		err := GlobalWebSocketCon.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			log.Println("Failed to send message to WebSocket:", err)
		}
	}
}
