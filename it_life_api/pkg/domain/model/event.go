package model

import "time"

type (
	Event struct {
		ID                uint64    `json:"id"`
		UserID            uint64    `json:"userId"`
		Name              string    `json:"name"`
		Description       string    `json:"description"`
		BeginningWeekDate time.Time `json:"beginningWeekDate"`
		StartsAt          time.Time `json:"startsAt"`
		EndsAt            time.Time `json:"endsAt"`
	}
)
