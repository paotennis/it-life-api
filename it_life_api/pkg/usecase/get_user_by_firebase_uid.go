package usecase

import (
	"github.com/rikuhatano09/it-life-api/internal/infrastructure/persistence"
	"github.com/rikuhatano09/it-life-api/pkg/domain/model"
)

func GetUserByFirebaseUID(firebaseUID string) (model.User, error) {
	userPersistence := persistence.NewUserPersistence()

	user, err := userPersistence.FindByFirebaseUID(firebaseUID)
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}
