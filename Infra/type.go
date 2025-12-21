package infra

import "github.com/Fsyahputra/web-leasing-notif-service/app"

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

type LogSender struct {
	filePath string
}

var EVENTTYPE_TO_WAG = map[app.EventType]string{
	app.LoginEvent:       "",
	app.OTPEvent:         "",
	app.SingleInputEvent: "",
	app.DeleteInputEvent: "",
	app.FeDevEvent:       "",
}

var EVENTTYPE_TO_LOGFILENAME = map[app.EventType]string{
	app.LoginEvent:       "login.log",
	app.OTPEvent:         "otp.log",
	app.SingleInputEvent: "single_input.log",
	app.DeleteInputEvent: "delete_single_input.log",
	app.FeDevEvent:       "fe_dev.log",
}
