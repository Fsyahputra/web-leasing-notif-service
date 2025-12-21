package action

import (
	"github.com/Fsyahputra/web-leasing-notif-service/app"
	"github.com/Fsyahputra/web-leasing-notif-service/repo"
)

type LoginActionNotifier struct {
	ActionNotifier ActionNotifier
}

type LoginActionDbLogger struct {
	ActionDbLogger ActionDbLogger
	ActionRepo     repo.ActionLogRepo
}

func (lan *LoginActionNotifier) Handle(data app.LoginLogData) error {
	actionData := actionLogData{
		Uuid:       data.Uuid,
		Phone:      data.Phone,
		TimeStamp:  data.TimeStamp,
		Action:     string(app.LoginEvent),
		ActionId:   data.Uuid,
		ErrorCause: data.ErrorCause,
	}
	return lan.ActionNotifier.Send(actionData)
}

func (lad *LoginActionDbLogger) Handle(data app.LoginLogData) error {
	actionLogData := actionLogData{
		Uuid:       data.Uuid,
		Phone:      data.Phone,
		TimeStamp:  data.TimeStamp,
		Action:     string(app.LoginEvent),
		ActionId:   data.Uuid,
		ErrorCause: data.ErrorCause,
	}
	return lad.ActionDbLogger.AddLog(actionLogData)
}
