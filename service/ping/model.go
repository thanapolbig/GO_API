package ping

import "time"

type heartbeat struct {
	Message string `json:"message"`
	DateTime time.Time `json:"date_time"`
}
