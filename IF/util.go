package IF

import (
	"encoding/json"
	"log"

	"github.com/Fsyahputra/web-leasing-notif-service/app"
	"github.com/IBM/sarama"
)

func unmarshallJson[T any](data []byte) (*T, error) {
	var result T
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func handleAll[T any](handlers []app.EventHandler[T], data T, priorityIdx int) error {
	for i, handler := range handlers {
		err := handler.Handle(data)
		if err == nil {
			continue
		}
		if i == priorityIdx {
			return err
		}

	}
	return nil
}

func consumeClaim[T any](
	session sarama.ConsumerGroupSession,
	claim sarama.ConsumerGroupClaim,
	handlers []app.EventHandler[T],
	priorityIdx int,

) {
	for message := range claim.Messages() {
		logData, err := unmarshallJson[T](message.Value)
		if err != nil {
			log.Printf("Error unmarshalling log data: %v\n", err)
			continue
		}
		err = handleAll[T](handlers, *logData, priorityIdx)
		if err != nil {
			log.Printf("Error handling log data: %v\n", err)
			continue
		}
		session.MarkMessage(message, "")
	}

}
