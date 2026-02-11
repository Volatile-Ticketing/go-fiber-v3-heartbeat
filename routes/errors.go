package routes

import "time"

type ErrorLog struct {
	Time    time.Time `json:"time"`
	Message string    `json:"message"`
}
