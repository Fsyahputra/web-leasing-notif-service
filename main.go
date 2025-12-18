package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/IBM/sarama"
)

type ReqBody struct {
	Msg       string `json:"message"`
	SessionId string `json:"session_id"`
	To        string `json:"to"`
}

func configureReq(req *http.Request) {
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Tenant", "huntersmithnusantara")

}

func sendToWhatsapp(msg string) {
	apiUrl := "https://api-wa.onestopcheck.id/api/whatsapp/send/text-notif"
	sessionId := "7b1a8c7c-d107-44e4-abf3-6565371da4ca"
	reqBody := ReqBody{
		Msg:       msg,
		SessionId: sessionId,
		To:        "6282250712167@c.us",
	}
	jsonData, _ := json.Marshal(reqBody)
	client := &http.Client{}
	req, _ := http.NewRequest("POST", apiUrl, bytes.NewBuffer(jsonData))
	configureReq(req)
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	fmt.Println("Response Status:", resp.Status)

}

func main() {
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
		sendToWhatsapp(string(msg.Value))
	}
}
