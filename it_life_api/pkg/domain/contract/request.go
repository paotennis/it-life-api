package contract

import "time"

type (
	// UserPostRequestBody is request body for creating a user.
	UserPostRequestBody struct {
		FirebaseUID string `json:"firebaseUid"`
		Email       string `json:"email"`
		Age         uint32 `json:"age"`
		Role        string `json:"role"`
		Company     string `json:"company"`
	}
)

type (
	// EventPostRequestBody is request body for creating a event.
	EventPostRequestBody struct {
		UserID uint64      `json:"userId"`
		Date   string      `json:"date"`
		Events []EventItem `json:"events"`
	}
)

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
		// TODO: Custom error handling
		panic(err)
	}
	return startsAt
}

func (eventItem EventItem) GetEndsAt(date string) time.Time {
	endsAt, err := time.Parse("2006/01/02 15:04", date+" "+eventItem.EndTime)
	if err != nil {
		// TODO: Custom error handling
		panic(err)
	}
	return endsAt
}
