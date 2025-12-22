package wiring

import (
	"github.com/Fsyahputra/web-leasing-notif-service/Infra"
	"github.com/Fsyahputra/web-leasing-notif-service/app"
	"github.com/Fsyahputra/web-leasing-notif-service/app/Action"
)

var wagConf = infra.WagConf{
	Login: infra.LoginGroups{
		LogGroupID: "120363420513101473@g.us",
	},
	OTP: infra.OTPGroups{
		LogGroupID: "120363421369311585@g.us",
	},
	SingleInput: infra.SIGroups{
		LogGroupID: "120363421369311585@g.us",
	},
	DeleteSingleInput: infra.DSIGroups{
		LogGroupID: "120363406247806640@g.us",
	},
	FeDev: infra.FeDevGroups{
		LogGroupID:     "120363403873333764@g.us",
		CompanyGroupID: "120363327165003756@g.us",
	},
	UpdateSingleInput: infra.UpdateSIGroups{
		LogGroupID: "120363405934976050@g.us",
	},
	Action: infra.ActionGroups{
		LogGroupID: "120363405962710549@g.us",
	},
}

type ActionHandlers struct {
	ActionNotifier action.ActionNotifier
	ActionDbLogger action.ActionDbLogger
}

type LoginHandlers struct {
	ActionHandlers
	LoginNotifier   app.LoginHandlerNotifier
	LoginDbLogger   app.LoginDbLogger
	LoginFileLogger app.LoginFileLogger
}

type OTPHandlers struct {
	ActionHandlers
	OTPNotifier   app.OTPHandlerNotifier
	OTPDbLogger   app.OTPDbLogger
	OTPFileLogger app.OTPFileLogger
}

type SingleInputHandlers struct {
	ActionHandlers
	SingleInputNotifier   app.SingleInputNotifier
	SingleInputDbLogger   app.SingleInputDbLogger
	SingleInputFileLogger app.SingleInputFileLogger
}

type UpdateSingleInputHandlers struct {
	ActionHandlers
	UpdateSingleInputNotifier   app.UpdateSingleInputNotifier
	UpdateSingleInputDbLogger   app.UpdateSingleInputDbLogger
	UpdateSingleInputFileLogger app.UpdateSingleInputFileLogger
}

type DeleteSingleInputHandlers struct {
	ActionHandlers
	DeleteSingleInputNotifier   app.DeleteSingleInputNotifier
	DeleteSingleInputDbLogger   app.DeleteSingleInputDbLogger
	DeleteSingleInputFileLogger app.DeleteSingleInputFileLogger
}

type FeDevLogHandlers struct {
	FeDevLogCompany app.EventHandler[app.FeDevLogData]
	FeDevLogLogger  app.EventHandler[app.FeDevLogData]
}

type Handlers struct {
	// Login         LoginHandlers
	// OTP           OTPHandlers
	// SingleInput   SingleInputHandlers
	// UpdateSI      UpdateSingleInputHandlers
	// DeleteSI      DeleteSingleInputHandlers
	FeDevNotifier FeDevLogHandlers
}

type FeDevSenders struct {
	LogGroup app.Sender
	Company  app.Sender
}

type senders struct {
	Login       app.Sender
	OTP         app.Sender
	SingleInput app.Sender
	UpdateSI    app.Sender
	DeleteSI    app.Sender
	FeDev       FeDevSenders
}

func getSender() senders {
	LoginSender := infra.NewWhatsAppApiSender(wagConf.Login.LogGroupID)
	OTPSender := infra.NewWhatsAppApiSender(wagConf.OTP.LogGroupID)
	SingleInputSender := infra.NewWhatsAppApiSender(wagConf.SingleInput.LogGroupID)
	UpdateSISender := infra.NewWhatsAppApiSender(wagConf.UpdateSingleInput.LogGroupID)
	DeleteSISender := infra.NewWhatsAppApiSender(wagConf.DeleteSingleInput.LogGroupID)
	FeDevLogSender := infra.NewWhatsAppApiSender(wagConf.FeDev.LogGroupID)
	FeDevCompanySender := infra.NewWhatsAppApiSender(wagConf.FeDev.CompanyGroupID)
	return senders{
		Login:       LoginSender,
		OTP:         OTPSender,
		SingleInput: SingleInputSender,
		UpdateSI:    UpdateSISender,
		DeleteSI:    DeleteSISender,
		FeDev: FeDevSenders{
			LogGroup: FeDevLogSender,
			Company:  FeDevCompanySender,
		},
	}

}

func GetHandlers() Handlers {
	senders := getSender()
	return Handlers{
		FeDevNotifier: FeDevLogHandlers{
			FeDevLogCompany: app.NewFeDevLogNotifier(senders.FeDev.Company, ""),
			FeDevLogLogger:  app.NewFeDevLogger(senders.FeDev.LogGroup),
		},
	}
}
