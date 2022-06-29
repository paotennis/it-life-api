package contract

import "time"

type (
	// UserWithWeekDataResponse is response of user with week data.
	UserWithWeekDataResponse struct {
		UserID   uint64     `json:"userId"`
		Age      uint32     `json:"age"`
		Role     string     `json:"role"`
		Company  string     `json:"company"`
		WeekData []WeekItem `json:"weekData"`
	}

	// WeekItem is an element of week data.
	WeekItem struct {
		Date       string `json:"date"`
		DateCount  uint32 `json:"dateCount"`
		EventCount uint32 `json:"eventCount"`
	}
)

func (weekItem WeekItem) ParseDateStringToTime() time.Time {
	date, err := time.Parse("2006/01/02", weekItem.Date)
	if err != nil {
		// TODO: Custom error handling.
		panic(err)
	}
	return date
}
