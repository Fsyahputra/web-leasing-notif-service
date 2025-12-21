package action

import (
	"fmt"
	"strings"

	"github.com/Fsyahputra/web-leasing-notif-service/app"
	"github.com/Fsyahputra/web-leasing-notif-service/repo"
)

type ActionNotifier struct {
	Sender app.Sender
}

func notifFormatter(data actionLogData) string {
	header := app.MakeBold(fmt.Sprintf("Notifikasi Aksi Pengguna  %s 📢", strings.ToUpper(data.Action)))
	var actionId string = "-"
	if data.ActionId != "" {
		actionId = data.ActionId
	}
	message := fmt.Sprintf(
		`
		%s	
		------------------------
		Uuid Pengguna : %s
		Nomor Handphone : %s
		Waktu : %s
		Aksi : %s
		Aksi Id : %s
		------------------------
		`,
		header,
		data.Uuid,
		data.Phone,
		fmt.Sprintf("%v", data.TimeStamp),
		data.Action,
		actionId,
	)
	if data.ErrorCause != "" {
		message += fmt.Sprintf("\nPenyebab Error : %s\n\n", data.ErrorCause)
	}
	return message
}

func NewActionNotifier(sender app.Sender) *ActionNotifier {
	return &ActionNotifier{
		Sender: sender,
	}

}

func (ad *ActionNotifier) Send(data actionLogData) error {
	message := notifFormatter(data)
	return ad.Sender.Send(message)
}

type ActionDbLogger struct {
	Arp repo.ActionLogRepo
}

func (dsl *ActionDbLogger) AddLog(data actionLogData) error {
	actionLog := repo.ActionLogData{
		Uuid:       data.Uuid,
		Phone:      data.Phone,
		TimeStamp:  data.TimeStamp,
		Action:     data.Action,
		ActionId:   data.ActionId,
		ErrorCause: data.ErrorCause,
	}
	return dsl.Arp.AddActionLog(actionLog)
}

func NewActionDbLogger(arp repo.ActionLogRepo) *ActionDbLogger {
	return &ActionDbLogger{
		Arp: arp,
	}
}
