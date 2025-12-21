package app

import (
	"fmt"

	"github.com/Fsyahputra/web-leasing-notif-service/repo"
)

type DeleteSingleInputNotifier struct {
	Sender Sender
	Usrp   repo.UserRepo
}

func (dh *DeleteSingleInputNotifier) Handle(data SingleInputLogData) error {
	userName, err := dh.Usrp.GetUserName(data.Uuid)
	if err != nil {
		return err
	}
	header := MakeBold("Log Aktivitas Pengguna 📋")
	message := fmt.Sprintf(
		`
		%s
		------------------------
		Aksi : Hapus Data Kendaraan
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
	if err := dh.Sender.Send(message); err != nil {
		return err
	}
	return nil

}

type DeleteSingleInputDbLogger struct {
	Vrp  repo.VehicleLogRepo
	Usrp repo.UserRepo
}

func (dl *DeleteSingleInputDbLogger) Handle(data SingleInputLogData) error {
	vehicleLog := repo.VehicleLogData{
		Uuid:      data.Uuid,
		Nopol:     data.VehicleData.Nopol,
		Noka:      data.VehicleData.Noka,
		TimeStamp: data.TimeStamp,
		Action:    "DELETE",
	}
	if err := dl.Vrp.AddVehicleLog(vehicleLog); err != nil {
		return err
	}
	return nil
}

type DeleteSingleInputFileLogger struct {
	Sender Sender
}

func (df *DeleteSingleInputFileLogger) Handle(data SingleInputLogData) error {
	fmtMsg := fmt.Sprintf("TimeStamp: %v | Uuid: %s | Phone: %s | Nopol: %s | Noka: %s | Cabang: %s | ErrorCause: %s\n | Action: DELETE",
		data.TimeStamp,
		data.Uuid,
		data.Phone,
		data.VehicleData.Nopol,
		data.VehicleData.Noka,
		data.VehicleData.Cabang,
		data.ErrorCause,
	)
	return df.Sender.Send(fmtMsg)
}
