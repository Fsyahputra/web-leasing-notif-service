package app

import (
	"fmt"

	"github.com/Fsyahputra/web-leasing-notif-service/repo"
)

type WALoginHandler struct {
	Sender Sender
	Usrp   repo.UserRepo
}

func (wh *WALoginHandler) Handle(data LoginLogData) error {
	userName, err := wh.Usrp.GetUserName(data.Uuid)
	if err != nil {
		return err
	}
	header := makeBold("Login Berhasil ✅")
	if data.ErrorCause != "" {
		header = makeBold("Login Gagal ❌")
	}
	message := fmt.Sprintf(
		`
		%s
		------------------------
		Nama Pengguna : %s
		Nomor Handphone : %s
		Waktu : %s
		------------------------
		`,
		header,
		userName,
		data.Phone,
		fmt.Sprintf("%v", data.TimeStamp),
	)
	if data.ErrorCause != "" {
		message += fmt.Sprintf("\nPenyebab Error : %s", data.ErrorCause)
	}
	if err := wh.Sender.Send(message); err != nil {
		return err
	}
	return nil
}

type LoginLogger struct {
	Lrp repo.AuthLogRepo
}

func (ll *LoginLogger) Handle(data LoginLogData) error {
	var err bool
	err = false
	if data.ErrorCause != "" {
		err = true
	}
	repoData := repo.LoginLogData{
		Uuid:       data.Uuid,
		Phone:      data.Phone,
		TimeStamp:  data.TimeStamp,
		ErrorCause: data.ErrorCause,
		Action:     "LOGIN",
		Error:      err,
	}
	return ll.Lrp.AddLoginLog(repoData)
}

type LoginFileLogger struct {
	Sender Sender
}

func (l *LoginFileLogger) Handle(data LoginLogData) error {
	fmtMessage := fmt.Sprintf(
		`TimeStamp: %v | UUID: %s | Phone: %s | ErrorCause: %s`,
		data.TimeStamp,
		data.Uuid,
		data.Phone,
		data.ErrorCause,
	)
	return l.Sender.Send(fmtMessage)
}
