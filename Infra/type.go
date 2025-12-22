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

type LoginGroups struct {
	LogGroupID string
}

type OTPGroups struct {
	LogGroupID string
}

type SIGroups struct {
	LogGroupID string
}

type DSIGroups struct {
	LogGroupID string
}

type FeDevGroups struct {
	LogGroupID     string
	CompanyGroupID string
}

type UpdateSIGroups struct {
	LogGroupID string
}

type ActionGroups struct {
	LogGroupID string
}

type WagConf struct {
	Login             LoginGroups
	OTP               OTPGroups
	SingleInput       SIGroups
	DeleteSingleInput DSIGroups
	FeDev             FeDevGroups
	UpdateSingleInput UpdateSIGroups
	Action            ActionGroups
}

var EVENTTYPE_TO_LOGFILENAME = map[app.EventType]string{
	app.LoginEvent:       "login.log",
	app.OTPEvent:         "otp.log",
	app.SingleInputEvent: "single_input.log",
	app.DeleteInputEvent: "delete_single_input.log",
	app.FeDevEvent:       "fe_dev.log",
}
