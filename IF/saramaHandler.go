package IF

import (
	"encoding/json"
	"log"

	"github.com/Fsyahputra/web-leasing-notif-service/app"
	"github.com/IBM/sarama"
)

type FeDevLogSaramaHandler struct {
	handler app.FeDevLogDataHandler
}

func (sh *FeDevLogSaramaHandler) Setup(sarama.ConsumerGroupSession) error {
	return nil

}

func (sh *FeDevLogSaramaHandler) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

func (sh *FeDevLogSaramaHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {

	for message := range claim.Messages() {
		var logData app.FeDevLogData
		err := json.Unmarshal(message.Value, &logData)
		if err != nil {
			log.Printf("Error unmarshaling message: %v", err)
			continue
		}
		err = sh.handler.Handle(logData)
		if err != nil {
			log.Printf("Error handling log data: %v", err)
			continue
		}
		session.MarkMessage(message, "")

	}
	return nil
}

func NewFeDevLogSaramaHandler(handler app.FeDevLogDataHandler) sarama.ConsumerGroupHandler {
	return &FeDevLogSaramaHandler{handler: handler}
}
