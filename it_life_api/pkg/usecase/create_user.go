package usecase

import (
	"github.com/rikuhatano09/it-life-api/internal/infrastructure/persistence"
	"github.com/rikuhatano09/it-life-api/pkg/domain/contract"
	"github.com/rikuhatano09/it-life-api/pkg/domain/model"
)

func CreateUser(requestBody contract.UserPostRequestBody) (model.User, error) {
	userPersistence := persistence.NewUserPersistence()

	user := model.User{
		FirebaseUID: requestBody.FirebaseUID,
		Email:       requestBody.Email,
		Age:         requestBody.Age,
		Role:        requestBody.Role,
		Company:     requestBody.Company,
	}

	user, err := userPersistence.Create(user)
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}
