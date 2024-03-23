package models

import "time"

type GoodsEvent struct {
	ID          int
	ProjectID   int
	Name        string
	Description string
	Priority    int
	Removed     bool
	EventTime   time.Time
}
