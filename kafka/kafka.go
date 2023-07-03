package kafka

import (
	"context"
	"log"
	"os"
	"sync"
	"time"

	"github.com/Shopify/sarama"
)

func SetupKafkaProducer() (sarama.SyncProducer, error) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5
	config.Producer.Timeout = 5 * time.Second
	config.Producer.Return.Successes = true
	config.Consumer.Offsets.Initial = sarama.OffsetOldest

	producer, err := sarama.NewSyncProducer([]string{"localhost:9092"}, config)
	if err != nil {
		return nil, err
	}
	return producer, nil
}

func SendUserToKafka(producer sarama.SyncProducer, user []byte) error {
	// dataJSON, err := json.Marshal(user)
	// if err != nil {
	// 	log.Println("Error marshaling data kafka:", err)
	// }
	msg := &sarama.ProducerMessage{
		Topic:     "ahkamu",
		Value:     sarama.StringEncoder(user),
		Timestamp: time.Now(),
	}

	_, _, err := producer.SendMessage(msg)
	return err

}

func StartKafkaConsumer() error {
	// Create Kafka consumer
	config := sarama.NewConfig()
	config.Consumer.Offsets.Initial = sarama.OffsetOldest
	topic := os.Getenv("TOPIC_BROKER")
	kafkaHost := os.Getenv("KAFKA_HOST")
	kafkaGroupID := os.Getenv("KAFKA_GROUP_ID")

	consumer, err := sarama.NewClient([]string{kafkaHost}, nil)
	if err != nil {
		log.Fatalf("Failed to create Kafka consumer: %s", err)
	}
	defer consumer.Close()

	group, err := sarama.NewConsumerGroupFromClient(kafkaGroupID, consumer)
	if err != nil {
		return err
	}
	defer group.Close()

	go func() {
		for err := range group.Errors() {
			log.Printf("Consumer group error: %v", err)
		}
	}()
	// Start consuming messages

	handler := &ConsumerGroupHandler{}
	wg := &sync.WaitGroup{}
	wg.Add(1)

	ctx := context.Background()

	go func() {
		defer wg.Done()
		for {
			err := group.Consume(ctx, []string{topic}, handler)
			if err != nil {
				log.Printf("Error from consumer: %v", err)
			}
			// Exit the loop if the consumer group is closed or encounters an error
			if handler.Closed() {
				return
			}
		}
	}()

	wg.Wait()
	return nil
}
