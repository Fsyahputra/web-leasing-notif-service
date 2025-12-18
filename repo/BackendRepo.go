package repo

type UserRepo interface {
	GetUserName(uuid string) (string, error)
	GetLeasing(uuid string) (string, error)
}

type OtpRepo interface {
	GetOTP(uuid string) (string, error)
	GetOTPIdByUserUUID(uuid string) (string, error)
}
