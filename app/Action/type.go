package action

type actionLogData struct {
	Uuid       string `json:"uuid"`
	Phone      string `json:"phone"`
	Action     string `json:"action"`
	TimeStamp  int64  `json:"timestamp"`
	ActionId   string `json:"action_id"`
	ErrorCause string `json:"error_cause,omitempty"`
}
