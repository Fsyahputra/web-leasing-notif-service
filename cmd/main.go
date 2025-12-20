package main

import (
	"fmt"
	"log"

	"github.com/IBM/sarama"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	brokerList := []string{"localhost:9092"}
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true

	master, err := sarama.NewConsumer(brokerList, config)
	if err != nil {
		log.Fatalln("Failed to start consumer:", err)
	}
	defer master.Close()

	consumer, err := master.ConsumePartition("test-topic", 0, sarama.OffsetOldest)
	if err != nil {
		log.Fatalln("Failed to start partition consumer:", err)
	}
	defer consumer.Close()

	fmt.Println("Consuming messages from test-topic...")
	for msg := range consumer.Messages() {
		fmt.Printf("Message: %s\n", string(msg.Value))
	}
}
