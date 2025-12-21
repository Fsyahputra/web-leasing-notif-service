package action

import (
	"encoding/json"
	"github.com/Fsyahputra/web-leasing-notif-service/app"
	"github.com/Fsyahputra/web-leasing-notif-service/repo"
)

type UpdateSIActionNotifier struct {
	ActionNotifier ActionNotifier
}

type UpdateSIActionDbLogger struct {
	ActionDbLogger ActionDbLogger
	ActionRepo     repo.ActionLogRepo
}

func (uan *UpdateSIActionNotifier) Handle(data app.SingleInputLogData) error {
	return handleSingleInputNotifier(
		uan.ActionNotifier,
		data,
		app.UpdateSingleInputEvent,
	)
}

func (udn *UpdateSIActionDbLogger) parseActionId(ids []string) string {
	jsonIds, err := json.Marshal(ids)
	if err != nil {
		return ""
	}
	return string(jsonIds)
}

func (udn *UpdateSIActionDbLogger) Handle(data app.SingleInputLogData) error {
	return handleSingleInputDbLogger(
		udn.ActionDbLogger,
		data,
		app.UpdateSingleInputEvent,
		udn.parseActionId,
	)
}
