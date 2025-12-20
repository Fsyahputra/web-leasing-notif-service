package app

import (
	"fmt"

	"github.com/Fsyahputra/web-leasing-notif-service/repo"
)

type WADeleteSingleInputHandler struct {
	Sender Sender
	Usrp   repo.UserRepo
}

func (dh *WADeleteSingleInputHandler) Handle(data DeleteSingleInputLogData) error {
	userName, err := dh.Usrp.GetUserName(data.Uuid)
	if err != nil {
		return err
	}
	header := makeBold("Log Aktivitas Pengguna 📋")
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
	if err := dh.Sender.Send(message); err != nil {
		return err
	}
	return nil

}

type DeleteSingleInputLogger struct {
	Vrp  repo.VehicleLogRepo
	Usrp repo.UserRepo
}

func (dl *DeleteSingleInputLogger) Handle(data DeleteSingleInputLogData) error {
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
