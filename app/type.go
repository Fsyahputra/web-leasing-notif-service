package app

type EventHandler[T any] interface {
	Handle(data T) error
}

type Sender interface {
	Send(message string) error
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

type FeDevLogData struct {
	CommitMsg string `json:"commit_msg,omitempty"`
	Author    string `json:"author,omitempty"`
	TimeStamp string `json:"timestamp,omitempty"`
	Diff      string `json:"diff,omitempty"`
}

type FeDevLogDataRaw struct {
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

type EventType string

const (
	LoginEvent             EventType = "WEB_LEASING.LOGIN_EVENT"
	OTPEvent               EventType = "WEB_LEASING.OTP_EVENT"
	SingleInputEvent       EventType = "WEB_LEASING.SINGLE_INPUT_EVENT"
	DeleteInputEvent       EventType = "WEB_LEASING.DELETE_INPUT_EVENT"
	FeDevEvent             EventType = "WEB_LEASING.FE_DEV_EVENT"
	UpdateSingleInputEvent EventType = "WEB_LEASING.UPDATE_SINGLE_INPUT_EVENT"
	ActionEvent            EventType = "WEB_LEASING.ACTION_EVENT"
)
