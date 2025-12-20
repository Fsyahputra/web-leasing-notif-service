package app

import (
	"fmt"

	"github.com/Fsyahputra/web-leasing-notif-service/repo"
)

type WASingleInputHandler struct {
	Sender Sender
	Usrp   repo.UserRepo
}

func (sh *WASingleInputHandler) Handle(data SingleInputLogData) error {
	userName, err := sh.Usrp.GetUserName(data.Uuid)
	if err != nil {
		return err
	}
	header := makeBold("Log Aktivitas Pengguna 📋")
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

	header2 := makeBold("Data Kendaraan 🏍️/🚘")
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

type SingleInputLogger struct {
	Vrp  repo.VehicleLogRepo
	Usrp repo.UserRepo
}

func (sl *SingleInputLogger) Handle(data SingleInputLogData) error {
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
