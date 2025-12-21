package action

import "github.com/Fsyahputra/web-leasing-notif-service/app"

func handleSingleInputNotifier(
	notifier ActionNotifier,
	data app.SingleInputLogData,
	eventType app.EventType,

) error {
	actionId := ""
	ids := []string{data.VehicleData.Nopol, data.VehicleData.Noka, data.VehicleData.Nosin}
	for _, id := range ids {
		if id != "" {
			actionId = id
			break
		}
	}
	actionData := actionLogData{
		Uuid:       data.Uuid,
		Phone:      data.Phone,
		TimeStamp:  data.TimeStamp,
		Action:     string(eventType),
		ActionId:   actionId,
		ErrorCause: data.ErrorCause,
	}
	return notifier.Send(actionData)
}

func handleSingleInputDbLogger(
	logger ActionDbLogger,
	data app.SingleInputLogData,
	eventType app.EventType,
	parseActionId func([]string) string,
) error {
	ids := []string{data.VehicleData.Nopol, data.VehicleData.Noka, data.VehicleData.Nosin}
	actionId := parseActionId(ids)

	actionData := actionLogData{
		Uuid:       data.Uuid,
		Phone:      data.Phone,
		TimeStamp:  data.TimeStamp,
		Action:     string(eventType),
		ActionId:   actionId,
		ErrorCause: data.ErrorCause,
	}
	return logger.AddLog(actionData)
}
