package infra

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/Fsyahputra/web-leasing-notif-service/app"
)

func getWAConf() (string, string, string) {
	apiUrl := os.Getenv("WHATSAPP_API_URL")
	sessionId := os.Getenv("WHATSAPP_API_SESSION_ID")
	tenant := os.Getenv("WHATSAPP_API_TENANT")
	return apiUrl, sessionId, tenant
}

func (s *WhatsAppAPISender) Send(message string) error {
	client := http.Client{}
	trimmedMessage := strings.TrimSpace(message)
	body := whatsAppAPIReqBody{
		Msg:       trimmedMessage,
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

func NewWhatsAppApiSender(event app.EventType) (*WhatsAppAPISender, error) {
	if _, ok := EVENTTYPE_TO_WAG[event]; !ok {
		return nil, fmt.Errorf("unsupported event type: %s", event)
	}
	apiUrl, sessionId, tenant := getWAConf()
	return &WhatsAppAPISender{
		config: WhatsAppAPISenderConfig{
			ApiUrl:    apiUrl,
			SessionId: sessionId,
			To:        EVENTTYPE_TO_WAG[event],
			Tenant:    tenant,
		},
	}, nil
}

func (l *LogSender) Send(message string) error {
	file, err := os.OpenFile(
		l.filePath,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0644,
	)
	if err != nil {
		return err
	}
	defer file.Close()
	log.SetOutput(file)
	log.SetFlags(log.LstdFlags | log.Ltime | log.Lshortfile)
	log.Println(message)
	return nil
}

func NewLogSender(event app.EventType) (*LogSender, error) {
	fileName, ok := EVENTTYPE_TO_LOGFILENAME[event]
	if !ok {
		return nil, fmt.Errorf("unsupported event type: %s", event)
	}
	fileDir := os.Getenv("LOG_DIR")
	if fileDir == "" {
		return nil, fmt.Errorf("LOG_DIR environment variable is not set")
	}
	fullPath := path.Join(fileDir, fileName)
	return &LogSender{
		filePath: fullPath,
	}, nil

}
