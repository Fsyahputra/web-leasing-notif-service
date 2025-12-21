package repo

type LoginLogData struct {
	Uuid       string
	Phone      string
	TimeStamp  int64
	ErrorCause string
	Action     string
	Error      bool
}

type OTPLogData struct {
	Uuid       string
	Phone      string
	OTPId      string
	TimeStamp  int64
	ErrorCause string
	Action     string
	Error      bool
}

type VehicleLogData struct {
	Uuid      string
	Nopol     string
	Noka      string
	TimeStamp int64
	Action    string
	Leasing   string
	Cabang    string
}

type ActionLogData struct {
	Uuid       string
	Phone      string
	Action     string
	TimeStamp  int64
	ActionId   string `default:""`
	ErrorCause string `default:""`
}

type AuthLogRepo interface {
	AddLoginLog(logData LoginLogData) error
	AddOTPLog(logData OTPLogData) error
}

type VehicleLogRepo interface {
	AddVehicleLog(VehicleLogData) error
}

type ActionLogRepo interface {
	AddActionLog(logData ActionLogData) error
}
