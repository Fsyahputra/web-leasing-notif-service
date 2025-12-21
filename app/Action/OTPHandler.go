package action

import (
	"github.com/Fsyahputra/web-leasing-notif-service/app"
	"github.com/Fsyahputra/web-leasing-notif-service/repo"
)

type OTPActionNotifier struct {
	ActionNotifier ActionNotifier
}

type OTPActionDbLogger struct {
	ActionDbLogger ActionDbLogger
	ActionRepo     repo.ActionLogRepo
}

func (oan *OTPActionNotifier) Handle(data app.OTPLogData) error {
	actionData := actionLogData{
		Uuid:       data.Uuid,
		Phone:      data.Phone,
		TimeStamp:  data.TimeStamp,
		Action:     string(app.OTPEvent),
		ActionId:   data.Uuid,
		ErrorCause: data.ErrorCause,
	}
	return oan.ActionNotifier.Send(actionData)
}

func (oad *OTPActionDbLogger) Handle(data app.OTPLogData) error {
	actionLogData := actionLogData{
		Uuid:       data.Uuid,
		Phone:      data.Phone,
		TimeStamp:  data.TimeStamp,
		Action:     string(app.OTPEvent),
		ActionId:   data.Uuid,
		ErrorCause: data.ErrorCause,
	}
	return oad.ActionDbLogger.AddLog(actionLogData)
}
