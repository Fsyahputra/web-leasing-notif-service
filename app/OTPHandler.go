package app

import (
	"fmt"
	"github.com/Fsyahputra/web-leasing-notif-service/repo"
	"log"
)

type OTPHandlerNotifier struct {
	Sender Sender
	Usrp   repo.UserRepo
	Otrp   repo.OtpRepo
}

func (wo *OTPHandlerNotifier) Handle(data OTPLogData) error {
	userName, err := wo.Usrp.GetUserName(data.Uuid)
	validOTP, err := wo.Otrp.GetOTP(data.Uuid)
	if err != nil {
		return err
	}
	header := MakeBold("OTP Berhasil Diverifikasi ✅")
	if data.ErrorCause != "" {
		header = MakeBold("OTP Gagal Diverifikasi ❌")

	}
	message := fmt.Sprintf(
		`
		%s
		------------------------
		Nama Pengguna : %s
		Nomor Handphone : %s
		OTP Yang Diterima : %s
		Valid OTP : %s
		Waktu : %s
		------------------------
		`,
		header,
		userName,
		data.Phone,
		data.OTP,
		validOTP,
		fmt.Sprintf("%v", data.TimeStamp),
	)
	if data.ErrorCause != "" {
		message += fmt.Sprintf("\nPenyebab Error : %s", data.ErrorCause)
	}

	if err := wo.Sender.Send(message); err != nil {
		return err
	}
	return nil
}

type OTPDbLogger struct {
	Otrp repo.OtpRepo
	Lrp  repo.AuthLogRepo
}

func (ol *OTPDbLogger) Handle(data OTPLogData) error {
	var error bool
	error = false
	if data.ErrorCause != "" {
		error = true
	}
	otpId, err := ol.Otrp.GetOTPIdByUserUUID(data.Uuid)
	if err != nil {
		return err
	}
	repoData := repo.OTPLogData{
		Uuid:       data.Uuid,
		Phone:      data.Phone,
		OTPId:      otpId,
		TimeStamp:  data.TimeStamp,
		ErrorCause: data.ErrorCause,
		Action:     "OTP Verification",
		Error:      error,
	}
	return ol.Lrp.AddOTPLog(repoData)
}

type OTPFileLogger struct {
	Sender Sender
}

func (fl *OTPFileLogger) Handle(data OTPLogData) error {
	fmtMsg := fmt.Sprintf(
		`TimeStamp: %s | UUID: %s | Phone: %s | OTP: %s | ErrorCause: %s`,
		fmt.Sprintf("%v", data.TimeStamp),
		data.Uuid,
		data.Phone,
		data.OTP,
		data.ErrorCause,
	)

	if err := fl.Sender.Send(fmtMsg); err != nil {
		log.Printf("Failed to write OTP log to file: %v", err)
		return err
	}
	return nil
}
