package infra

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func (s *WhatsAppAPISender) Send(message string) error {
	client := http.Client{}
	body := whatsAppAPIReqBody{
		Msg:       message,
		SessionId: s.config.SessionId,
		To:        s.config.To,
	}
	data, err := json.Marshal(body)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", s.config.ApiUrl, bytes.NewBuffer(data))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Tenant", s.config.Tenant)
	_, err = client.Do(req)
	return err
}

func NewWhatsAppApiSender(config WhatsAppAPISenderConfig) *WhatsAppAPISender {
	return &WhatsAppAPISender{
		config: config,
	}

}
