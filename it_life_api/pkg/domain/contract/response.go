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

	// UserWithWeekEventResponse is response of user with week events.
	UserWithWeekEventResponse struct {
		UserID    uint64      `json:"userId"`
		StartDate string      `json:"startDate"`
		Age       uint32      `json:"age"`
		Role      string      `json:"role"`
		Company   string      `json:"company"`
		Monday    []EventItem `json:"monday"`
		Tuesday   []EventItem `json:"tuesday"`
		Wednesday []EventItem `json:"wednesday"`
		Thursday  []EventItem `json:"thursday"`
		Friday    []EventItem `json:"friday"`
		Saturday  []EventItem `json:"saturday"`
		Sunday    []EventItem `json:"sunday"`
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
