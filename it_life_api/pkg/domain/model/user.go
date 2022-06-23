package model

type (
	User struct {
		ID          uint64 `json:"id"`
		FirebaseUID string `json:"firebaseUid"`
		Email       string `json:"email"`
		Age         uint32 `json:"age"`
		Role        string `json:"role"`
		Company     string `json:"company"`
	}
)
