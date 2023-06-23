package main

import (
	"encoding/json"
	"golang-kafka-sarama-gorm/db"
	kafkaGo "golang-kafka-sarama-gorm/kafka"
	"golang-kafka-sarama-gorm/models"
	"log"
	"net/http"
	"time"

	"github.com/bradfitz/gomemcache/memcache"
	"github.com/gin-gonic/gin"
)

func main() {
	// Set up Gin router
	router := gin.Default()
	producer, err := kafkaGo.SetupKafkaProducer()
	if err != nil {
		log.Println("Failed to setup kafka:", err)
		return
	}
	db.ConnectDatabase()
	db.ConnectMemcached()
	// Define the route handler for creating a new user
	router.POST("/tracking", func(c *gin.Context) {
		var tracking models.Tracking

		// // Bind JSON request body to the user struct
		if err := c.ShouldBindJSON(&tracking); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// // Create the user in the database
		// db.DB.Create(&user)

		// Send user data to Kafka

		// tracking.CreatedAt = time.Now()

		timeStamps := time.Now()
		data := []interface{}{
			tracking,
			timeStamps,
		}
		dataJSON, errs := json.Marshal(data)
		if errs != nil {
			log.Println("Error marshaling data for Memcached:", err)
		} else {
			err = db.MC.Set(&memcache.Item{Key: tracking.TrackNumber, Value: dataJSON})
			if err != nil {
				log.Println("Error storing data in Memcached:", err)
			}
			err := kafkaGo.SendUserToKafka(producer, dataJSON)
			if err != nil {
				log.Print("Failed to produce message:")
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to produce message"})
				return
			}
		}

		c.JSON(http.StatusOK, tracking)
	})

	// Start the Gin server
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %s", err)
	}
}
