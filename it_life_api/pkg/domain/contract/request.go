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

	// EventPostRequestBody is request body for creating a event.
	EventPostRequestBody struct {
		UserID uint64      `json:"userId"`
		Date   string      `json:"date"`
		Events []EventItem `json:"events"`
	}
)

func (eventPostRequestBody EventPostRequestBody) GetBeginningWeekDate() time.Time {
	eventDate, err := time.Parse("2006/01/02", eventPostRequestBody.Date)
	if err != nil {
		// TODO: Custom error handling.
		panic(err)
	}
	weekDay := eventDate.Weekday()
	var beginningWeekDate time.Time

	if weekDay == time.Sunday {
		beginningWeekDate = eventDate.AddDate(0, 0, -6)
	} else {
		beginningWeekDate = eventDate.AddDate(0, 0, -int(weekDay)+1)
	}
	return beginningWeekDate
}
