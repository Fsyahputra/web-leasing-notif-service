package action

import (
	"github.com/Fsyahputra/web-leasing-notif-service/app"
	"github.com/Fsyahputra/web-leasing-notif-service/repo"
)

type SingleInputActionNotifier struct {
	ActionNotifier ActionNotifier
}

type SingleInputActionDbLogger struct {
	ActionDbLogger ActionDbLogger
	ActionRepo     repo.ActionLogRepo
}

func (sin *SingleInputActionNotifier) Handle(data app.SingleInputLogData) error {
	return handleSingleInputNotifier(
		sin.ActionNotifier,
		data,
		app.SingleInputEvent,
	)
}

func (sid *SingleInputActionDbLogger) parseActionId(ids []string) string {
	actionId := ""
	for _, id := range ids {
		if id != "" {
			actionId = id
			break
		}
	}
	return actionId
}

func (sid *SingleInputActionDbLogger) Handle(data app.SingleInputLogData) error {
	return handleSingleInputDbLogger(
		sid.ActionDbLogger,
		data,
		app.SingleInputEvent,
		sid.parseActionId,
	)
}
