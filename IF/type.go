package IF

import "github.com/Fsyahputra/web-leasing-notif-service/app"

var EVENTTYPE_TO_KAFKA_TOPICS = map[app.EventType]string{
	app.LoginEvent:       "web-leasing-login-event",
	app.OTPEvent:         "web-leasing-otp-event",
	app.SingleInputEvent: "single-input-data-event",
	app.DeleteInputEvent: "delete-single-input-event",
	app.FeDevEvent:       "fe-dev-log",
}
