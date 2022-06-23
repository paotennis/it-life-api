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
