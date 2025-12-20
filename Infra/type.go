package infra

type WhatsAppAPISenderConfig struct {
	ApiUrl    string
	SessionId string
	To        string
	Tenant    string
}

type whatsAppAPIReqBody struct {
	Msg       string `json:"message"`
	SessionId string `json:"session_id"`
	To        string `json:"to"`
}

type WhatsAppAPISender struct {
	config WhatsAppAPISenderConfig
}
