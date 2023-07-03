package main

import (
	"fmt"
	"log"

	kafkaGo "golang-kafka-sarama-gorm/kafka"
	websocketGo "golang-kafka-sarama-gorm/websocket"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println(err)
	}
	// Set up Gin router
	router := gin.Default()

	// Start the Kafka consumer in the background
	go func() {
		err := kafkaGo.StartKafkaConsumer()
		if err != nil {
			log.Fatalf("Consumer error: %v", err)
		}
	}()

	// Define a route handler for handling HTTP requests

	router.GET("/webSocket", func(c *gin.Context) {
		conn, err := websocketGo.Upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			log.Println("Failed to upgrade WebSocket connection:", err)
			return
		}
		defer conn.Close()

		websocketGo.GlobalWebSocketCon = conn // Store the connection globally

		for {
			// Read message from WebSocket
			_, message, err := conn.ReadMessage()
			if err != nil {
				log.Println("Failed to read message from WebSocket:", err)
				break
			}

			log.Println("Received message from WebSocket:", string(message))

			// Send a response back to the WebSocket client
			response := []byte("Received your message: " + string(message))
			err = conn.WriteMessage(websocket.TextMessage, response)
			if err != nil {
				log.Println("Failed to send response to WebSocket:", err)
				break
			}
		}
	})

	// Start the Gin server
	if err := router.Run(":8081"); err != nil {
		log.Fatalf("Failed to start server: %s", err)
	}

}
