package action

import (
	"encoding/json"

	"github.com/Fsyahputra/web-leasing-notif-service/app"
	"github.com/Fsyahputra/web-leasing-notif-service/repo"
)

type DeleteSIActionNotifier struct {
	ActionNotifier ActionNotifier
}

type DeleteSIActionDbLogger struct {
	ActionDbLogger ActionDbLogger
	ActionRepo     repo.ActionLogRepo
}

func (dan *DeleteSIActionNotifier) Handle(data app.SingleInputLogData) error {
	return handleSingleInputNotifier(
		dan.ActionNotifier,
		data,
		app.DeleteInputEvent,
	)
}

func (dbn *DeleteSIActionDbLogger) parseActionId(ids []string) string {
	jsonIds, err := json.Marshal(ids)
	if err != nil {
		return ""
	}
	return string(jsonIds)
}

func (dbn *DeleteSIActionDbLogger) Handle(data app.SingleInputLogData) error {
	return handleSingleInputDbLogger(
		dbn.ActionDbLogger,
		data,
		app.DeleteInputEvent,
		dbn.parseActionId,
	)
}
