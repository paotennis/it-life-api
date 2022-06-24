package model

import "time"

type (
	Event struct {
		ID          int64     `json:"id"`
		UserID      int64     `json:"userId"`
		Name        string    `json:"name"`
		Description string    `json:"description"`
		StartsAt    time.Time `json:"startsAt"`
		EndsAt      time.Time `json:"endsAt"`
	}
)
