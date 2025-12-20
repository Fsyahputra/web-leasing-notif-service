package app

type EventHandler[T any] interface {
	Handle(data T) error
}

type Sender interface {
	Send(message string) error
}

type WhatsAppAPISenderConfig struct {
	ApiUrl    string
	SessionId string
	To        string
	Tenant    string
}

type LoginLogData struct {
	Uuid       string `json:"uuid"`
	Phone      string `json:"phone"`
	TimeStamp  int64  `json:"timestamp"`
	ErrorCause string `json:"errorCause,omitempty"`
}

type OTPLogData struct {
	Uuid       string `json:"uuid"`
	Phone      string `json:"phone"`
	OTP        string `json:"otp"`
	TimeStamp  int64  `json:"timestamp"`
	ErrorCause string `json:"errorCause,omitempty"`
}

type VehicleLogData struct {
	Nopol  string `json:"nopol,omitempty"`
	Noka   string `json:"noka,omitempty"`
	Nosin  string `json:"nosin,omitempty"`
	Cabang string `json:"cabang,omitempty"`
}

type SingleInputLogData struct {
	Uuid        string         `json:"uuid"`
	Phone       string         `json:"phone"`
	TimeStamp   int64          `json:"timestamp"`
	ErrorCause  string         `json:"errorCause,omitempty"`
	VehicleData VehicleLogData `json:"vehicleData"`
}

type DeleteSingleInputLogData struct {
	Uuid        string         `json:"uuid"`
	Phone       string         `json:"phone"`
	TimeStamp   int64          `json:"timestamp"`
	ErrorCause  string         `json:"errorCause,omitempty"`
	VehicleData VehicleLogData `json:"vehicleData"`
}

type FeDevLogData struct {
	CommitMsg string `json:"commit_msg,omitempty"`
	Author    string `json:"author,omitempty"`
	TimeStamp string `json:"timestamp,omitempty"`
}

type FeDevLogDataProcessed struct {
	CommitMsg string
	Author    string
	TimeStamp string
	Summary   string
}
