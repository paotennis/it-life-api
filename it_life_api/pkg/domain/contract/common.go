package contract

import "time"

type (
	// EventItem is an element of events.
	EventItem struct {
		Name        string `json:"name"`
		StartTime   string `json:"startTime"`
		EndTime     string `json:"endTime"`
		Description string `json:"description"`
	}
)

func (eventItem EventItem) GetStartsAt(date string) time.Time {
	startsAt, err := time.Parse("2006/01/02 15:04", date+" "+eventItem.StartTime)
	if err != nil {
		// TODO: Custom error handling.
		panic(err)
	}
	return startsAt
}

func (eventItem EventItem) GetEndsAt(date string) time.Time {
	endsAt, err := time.Parse("2006/01/02 15:04", date+" "+eventItem.EndTime)
	if err != nil {
		// TODO: Custom error handling.
		panic(err)
	}
	return endsAt
}
