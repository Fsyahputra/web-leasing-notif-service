package IF

import (
	"github.com/Fsyahputra/web-leasing-notif-service/app"
	"github.com/IBM/sarama"
)

type noOpSaramaHandler struct{}

func (sh *noOpSaramaHandler) Setup(sarama.ConsumerGroupSession) error {
	return nil
}

func (sh *noOpSaramaHandler) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

type FeDevLogSaramaHandler struct {
	noOpSaramaHandler
	Handlers           []app.EventHandler[app.FeDevLogData]
	HandlerPriorityIdx int
}

type LoginLogSaramaHandler struct {
	noOpSaramaHandler
	Handlers           []app.EventHandler[app.LoginLogData]
	HandlerPriorityIdx int
}

type SingleInputLogDataSaramaHandler struct {
	noOpSaramaHandler
	Handlers           []app.EventHandler[app.SingleInputLogData]
	HandlerPriorityIdx int
}

type DeleteSingleInputLogDataSaramaHandler struct {
	noOpSaramaHandler
	Handlers           []app.EventHandler[app.DeleteSingleInputLogData]
	HandlerPriorityIdx int
}

type OTPLogDataHandler struct {
	noOpSaramaHandler
	Handlers           []app.EventHandler[app.OTPLogData]
	HandlerPriorityIdx int
}

func NewSingleInputLogSaramaHandler(handlers []app.EventHandler[app.SingleInputLogData], priorityIdx int) sarama.ConsumerGroupHandler {
	return &SingleInputLogDataSaramaHandler{
		noOpSaramaHandler:  noOpSaramaHandler{},
		Handlers:           handlers,
		HandlerPriorityIdx: priorityIdx,
	}
}

func (sh *SingleInputLogDataSaramaHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	consumeClaim[app.SingleInputLogData](session, claim, sh.Handlers, sh.HandlerPriorityIdx)
	return nil
}

func NewDeleteSingleInputLogSaramaHandler(handlers []app.EventHandler[app.DeleteSingleInputLogData], priorityIdx int) sarama.ConsumerGroupHandler {
	return &DeleteSingleInputLogDataSaramaHandler{
		noOpSaramaHandler:  noOpSaramaHandler{},
		Handlers:           handlers,
		HandlerPriorityIdx: priorityIdx,
	}
}

func (sh *DeleteSingleInputLogDataSaramaHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	consumeClaim[app.DeleteSingleInputLogData](session, claim, sh.Handlers, sh.HandlerPriorityIdx)
	return nil
}

func NewFeDevLogSaramaHandler(handlers []app.EventHandler[app.FeDevLogData], priorityIdx int) sarama.ConsumerGroupHandler {
	return &FeDevLogSaramaHandler{
		noOpSaramaHandler:  noOpSaramaHandler{},
		Handlers:           handlers,
		HandlerPriorityIdx: priorityIdx,
	}
}

func (sh *FeDevLogSaramaHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	consumeClaim[app.FeDevLogData](session, claim, sh.Handlers, sh.HandlerPriorityIdx)
	return nil
}

func NewLoginLogSaramaHandler(handlers []app.EventHandler[app.LoginLogData], priorityIdx int) sarama.ConsumerGroupHandler {
	return &LoginLogSaramaHandler{
		noOpSaramaHandler:  noOpSaramaHandler{},
		Handlers:           handlers,
		HandlerPriorityIdx: priorityIdx,
	}
}

func (sh *LoginLogSaramaHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	consumeClaim[app.LoginLogData](session, claim, sh.Handlers, sh.HandlerPriorityIdx)
	return nil
}

func OTPLogDataSaramaHandler(handlers []app.EventHandler[app.OTPLogData], priorityIdx int) sarama.ConsumerGroupHandler {
	return &OTPLogDataHandler{
		noOpSaramaHandler:  noOpSaramaHandler{},
		Handlers:           handlers,
		HandlerPriorityIdx: priorityIdx,
	}
}

func (sh *OTPLogDataHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	consumeClaim[app.OTPLogData](session, claim, sh.Handlers, sh.HandlerPriorityIdx)
	return nil
}
