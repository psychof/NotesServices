package domain

import "time"

type Notes struct {
	Title      string     `json:"Title"`
	Text       string     `json:"Text"`
	Time_stamp *time.Time `json:"timeStemp"`
}
