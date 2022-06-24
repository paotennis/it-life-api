package contract

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
		UserId int64       `json:"userId"`
		Date   string      `json:"date"`
		Events []EventItem `json:"events"`
	}
)

type (
	// EventItem is event structure definition.
	EventItem struct {
		Name        string `json:"name"`
		StartTime   string `json:"startTime"`
		EndTime     string `json:"endTime"`
		Description string `json:"description"`
	}
)
