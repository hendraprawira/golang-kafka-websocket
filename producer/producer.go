package main

import (
	"encoding/json"
	kafkaGo "golang-kafka-sarama-gorm/kafka"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MessageTest struct {
	MessageTest string `json:"message"`
}

func main() {
	// Set up Gin router
	router := gin.Default()
	producer, err := kafkaGo.SetupKafkaProducer()
	if err != nil {
		log.Println("Failed to setup kafka:", err)
		return
	}
	// Define the route handler for creating a new user
	router.GET("/send-data", func(c *gin.Context) {
		var messageTest MessageTest

		// // Bind JSON request body to the user struct
		if err := c.ShouldBindJSON(&messageTest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		dataJSON, errs := json.Marshal(messageTest)
		if errs != nil {
			log.Println("Error marshaling data for Memcached:", err)
		} else {
			err := kafkaGo.SendUserToKafka(producer, dataJSON)
			if err != nil {
				log.Print("Failed to produce message:")
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to produce message"})
				return
			}
		}

		c.JSON(http.StatusOK, messageTest)
	})

	// Start the Gin server
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %s", err)
	}
}
