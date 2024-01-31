package models

import "time"

var ScheduleData []Schedule

type Schedule struct {
	EditionNumber  string    `json:"edition_number"`
	UserID         int       `json:"user_id"`
	PickUpSchedule time.Time `json:"pick_up_schedule"`
}

type ScheduleRequest struct {
	EditionNumber  string    `json:"edition_number"`
	UserID         int       `json:"user_id"`
	PickUpSchedule time.Time `json:"pick_up_schedule"`
}
