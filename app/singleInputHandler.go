package app

import (
	"fmt"
	"github.com/Fsyahputra/web-leasing-notif-service/repo"
)

type SingleInputNotifier struct {
	Sender Sender
	Usrp   repo.UserRepo
}

func (sh *SingleInputNotifier) Handle(data SingleInputLogData) error {
	userName, err := sh.Usrp.GetUserName(data.Uuid)
	if err != nil {
		return err
	}
	header := MakeBold("Log Aktivitas Pengguna 📋")
	message := fmt.Sprintf(
		`
		%s
		------------------------
		Aksi : Input Data Kendaraan
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
		message += fmt.Sprintf("\nPenyebab Error : %s\n\n", data.ErrorCause)
	}

	header2 := MakeBold("Data Kendaraan 🏍️/🚘")
	message += fmt.Sprintf("\n%s\n", header2)
	message += fmt.Sprintf(
		`
			------------------------
			Data Kendaraan:
			Nopol : %s
			Noka : %s
			Cabang : %s
			------------------------
			`,
		data.VehicleData.Nopol,
		data.VehicleData.Noka,
		data.VehicleData.Cabang,
	)
	if err := sh.Sender.Send(message); err != nil {
		return err
	}
	return nil

}

func NewSingleInputNotifier(userRepo repo.UserRepo, sender Sender) *SingleInputNotifier {
	return &SingleInputNotifier{
		Sender: sender,
		Usrp:   userRepo,
	}
}

type SingleInputDbLogger struct {
	Vrp  repo.VehicleLogRepo
	Usrp repo.UserRepo
}

func (sl *SingleInputDbLogger) Handle(data SingleInputLogData) error {
	leasing, err := sl.Usrp.GetLeasing(data.Uuid)
	if err != nil {
		return err
	}
	repoData := repo.VehicleLogData{
		Uuid:      data.Uuid,
		Nopol:     data.Phone,
		Noka:      data.VehicleData.Noka,
		Leasing:   leasing,
		Cabang:    data.VehicleData.Cabang,
		Action:    "ADD",
		TimeStamp: data.TimeStamp,
	}
	return sl.Vrp.AddVehicleLog(repoData)
}

func NewSingleInputDbLogger(vrp repo.VehicleLogRepo, usrp repo.UserRepo) *SingleInputDbLogger {
	return &SingleInputDbLogger{
		Vrp:  vrp,
		Usrp: usrp,
	}

}

type SingleInputFileLogger struct {
	Sender Sender
}

func (sl *SingleInputFileLogger) Handle(data SingleInputLogData) error {
	fmtMsg := fmt.Sprintf(
		`TimeStamp: %s | UUID: %s | Phone: %s | Nopol: %s | Noka: %s | Cabang: %s | ErrorCause: %s | Action: INPUT`,
		fmt.Sprintf("%v", data.TimeStamp),
		data.Uuid,
		data.Phone,
		data.VehicleData.Nopol,
		data.VehicleData.Noka,
		data.VehicleData.Cabang,
		data.ErrorCause,
	)
	if err := sl.Sender.Send(fmtMsg); err != nil {
		return err
	}
	return nil
}

func NewSingleInputFileLogger(sender Sender) *SingleInputFileLogger {
	return &SingleInputFileLogger{
		Sender: sender,
	}
}
