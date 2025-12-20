package app

import (
	"fmt"
	"github.com/Fsyahputra/web-leasing-notif-service/repo"
)

type WAOTPHandler struct {
	Sender Sender[string]
	Usrp   repo.UserRepo
	Otrp   repo.OtpRepo
}

func (wo *WAOTPHandler) Handle(data OTPLogData) error {
	userName, err := wo.Usrp.GetUserName(data.Uuid)
	validOTP, err := wo.Otrp.GetOTP(data.Uuid)
	if err != nil {
		return err
	}
	header := makeBold("OTP Berhasil Diverifikasi ✅")
	if data.ErrorCause != "" {
		header = makeBold("OTP Gagal Diverifikasi ❌")

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

type OTPLogger struct {
	Otrp repo.OtpRepo
	Lrp  repo.AuthLogRepo
}

func (ol *OTPLogger) Handle(data OTPLogData) error {
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
