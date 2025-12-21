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
	app.LoginEvent:             "120363420513101473@g.us",
	app.OTPEvent:               "120363421369311585@g.us",
	app.SingleInputEvent:       "120363421369311585@g.us",
	app.DeleteInputEvent:       "120363406247806640@g.us",
	app.FeDevEvent:             "120363403873333764@g.us",
	app.UpdateSingleInputEvent: "120363405934976050@g.us",
	app.ActionEvent:            "120363405962710549@g.us",
}

var EVENTTYPE_TO_LOGFILENAME = map[app.EventType]string{
	app.LoginEvent:       "login.log",
	app.OTPEvent:         "otp.log",
	app.SingleInputEvent: "single_input.log",
	app.DeleteInputEvent: "delete_single_input.log",
	app.FeDevEvent:       "fe_dev.log",
}
