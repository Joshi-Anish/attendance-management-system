package models

import "time"

type Attendance struct {
	StudentID string    `json:"name"`
	CheckIn   time.Time `json:"check_in"`
	CheckOut  time.Time `json:"check_out"`
}
