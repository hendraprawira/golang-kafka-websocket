package main

import (
	"log"
	"net/http"

	"golang-kafka-sarama-gorm/db"
	kafkaGo "golang-kafka-sarama-gorm/kafka"
	"golang-kafka-sarama-gorm/models"

	"github.com/gin-gonic/gin"
)

func main() {
	// Set up Gin router
	db.ConnectMemcached()
	router := gin.Default()

	// Start the Kafka consumer in the background
	go func() {
		err := kafkaGo.StartKafkaConsumer()
		if err != nil {
			log.Fatalf("Consumer error: %v", err)
		}
	}()

	// Define a route handler for handling HTTP requests
	router.GET("/consume", func(c *gin.Context) {
		// Perform any desired actions when a request is received
		var data models.Tracking

		// Try to fetch data from Memcached
		item, err := db.MC.Get("1")
		if err == nil {
			// Data found in Memcached, unmarshal it
			log.Print(string(item.Value))
			// err = json.Unmarshal(item.Value, &data)
			// if err != nil {
			// 	log.Println("Error unmarshaling data from Memcached:", err)
			// 	c.JSON(http.StatusBadRequest, data)
			// }
		} else {
			log.Println("Error fetching data from Memcached:", err)
			c.JSON(http.StatusBadRequest, data)
		}
		c.JSON(http.StatusOK, data.AccelarationX)
	})

	// Start the Gin server
	if err := router.Run(":8000"); err != nil {
		log.Fatalf("Failed to start server: %s", err)
	}

}
