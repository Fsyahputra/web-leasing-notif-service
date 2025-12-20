package app

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type WhatsAppAPISender struct {
	config WhatsAppAPISenderConfig
}

type whatsAppAPIReqBody struct {
	Msg       string `json:"message"`
	SessionId string `json:"session_id"`
	To        string `json:"to"`
}

func (s *WhatsAppAPISenderConfig) Send(message string) error {
	client := http.Client{}
	body := whatsAppAPIReqBody{
		Msg:       message,
		SessionId: s.SessionId,
		To:        s.To,
	}
	data, err := json.Marshal(body)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", s.ApiUrl, bytes.NewBuffer(data))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Tenant", s.Tenant)
	_, err = client.Do(req)
	return err
}
