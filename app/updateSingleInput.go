package app

import (
	"fmt"
	"github.com/Fsyahputra/web-leasing-notif-service/repo"
)

type UpdateSingleInputNotifier struct {
	Sender Sender
	Usrp   repo.UserRepo
}

func (un *UpdateSingleInputNotifier) Handle(data SingleInputLogData) error {
	userName, err := un.Usrp.GetUserName(data.Uuid)
	if err != nil {
		return err
	}
	header := MakeBold("Log Aktivitas Pengguna 📋")
	message := fmt.Sprintf(
		`
		%s
		------------------------
		Aksi : Update Data Kendaraan
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
	if err := un.Sender.Send(message); err != nil {
		return err
	}
	return nil
}

type UpdateSingleInputDbLogger struct {
	Vrp  repo.VehicleLogRepo
	Usrp repo.UserRepo
}

func (ul *UpdateSingleInputDbLogger) Handle(data SingleInputLogData) error {
	leasing, err := ul.Usrp.GetLeasing(data.Uuid)
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
	if dbErr := ul.Vrp.AddVehicleLog(repoData); dbErr != nil {
		return dbErr
	}
	return nil
}

type UpdateSingleInputFileLogger struct {
	Sender Sender
}

func (uf *UpdateSingleInputFileLogger) Handle(data SingleInputLogData) error {
	fmtMsg := fmt.Sprintf(
		`TimeStamp: %s | UUID: %s | Phone: %s | Nopol: %s | Noka: %s | Cabang: %s | ErrorCause: %s | Action: UPDATE`,
		fmt.Sprintf("%v", data.TimeStamp),
		data.Uuid,
		data.Phone,
		data.VehicleData.Nopol,
		data.VehicleData.Noka,
		data.VehicleData.Cabang,
		data.ErrorCause,
	)
	if err := uf.Sender.Send(fmtMsg); err != nil {
		return err
	}
	return nil
}
