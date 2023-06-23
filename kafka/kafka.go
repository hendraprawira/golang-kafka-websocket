package kafka

import (
	"context"
	"log"
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
		Topic:     "ownship",
		Value:     sarama.StringEncoder(user),
		Timestamp: time.Now(),
	}

	_, _, err := producer.SendMessage(msg)
	return err

}

func StartKafkaConsumer() error {
	// Create Kafka consumer

	consumer, err := sarama.NewClient([]string{"localhost:9092"}, nil)
	if err != nil {
		log.Fatalf("Failed to create Kafka consumer: %s", err)
	}
	defer consumer.Close()

	group, err := sarama.NewConsumerGroupFromClient("your", consumer)
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
			err := group.Consume(ctx, []string{"ownship"}, handler)
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
