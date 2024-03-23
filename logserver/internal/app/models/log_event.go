package models

import "time"

type LogEvent struct {
	Time  time.Time
	Msg   string
	Level string
}
