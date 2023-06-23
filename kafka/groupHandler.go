package kafka

import (
	"fmt"
	"golang-kafka-sarama-gorm/db"
	"log"
	"time"

	"github.com/Shopify/sarama"
	"github.com/bradfitz/gomemcache/memcache"
)

type ConsumerGroupHandler struct {
	closed  bool
	message string
}

func (h *ConsumerGroupHandler) Setup(sarama.ConsumerGroupSession) error   { return nil }
func (h *ConsumerGroupHandler) Cleanup(sarama.ConsumerGroupSession) error { return nil }
func (h *ConsumerGroupHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	// var datas models.Tracking
	// timeStamps := time.Now()
	// var data models.Tracking

	for message := range claim.Messages() {
		start := time.Now()
		// Process the consumed message as desired
		// In this example, we store the message in the handler for retrieval in the route handler
		str := string(message.Value) // Convert byte array to string

		// Print the received message and duration

		// If the byte array is not encoded in UTF-8, you can use the following:
		// str := string(message[:])
		err := db.MC.Set(&memcache.Item{Key: "1", Value: message.Value})
		if err != nil {
			log.Println("Error storing data in Memcached:", err)
		}
		fmt.Println(str) // Output: "abcd"
		// fmt.Printf("Time taken: %s\n", duration)
		// err := json.Unmarshal(message.Value, &data)
		// if err != nil {
		// 	log.Fatalf("Failed to marshal array to JSON: %v", err)
		// }
		// log.Print(data)
		// date, _ := time.Parse(time.RFC3339, data[1].(string))
		// log.Print("date ", date)
		// processing := time.Since(date)
		// log.Print(data[1])
		// timeStamps1 := time.Now()
		// log.Print(timeStamps1)
		// log.Printf("Waktu %s", processing)
		// log.Print(data[1])
		// lol := message.Timestamp
		// processings := time.Since(lol)
		// log.Printf("Waktu hmm %s", processings)

		session.MarkMessage(message, "hmm")
		processing := time.Since(start)
		fmt.Printf("Time taken: %s\n", processing)
	}

	return nil
}

func (h *ConsumerGroupHandler) Closed() bool {
	return h.closed
}
