package persistence

import (
	"gorm.io/gorm"

	"github.com/rikuhatano09/it-life-api/pkg/domain/model"
	"github.com/rikuhatano09/it-life-api/pkg/domain/repository"
)

type (
	// UserPersistence is a persistence to preserve users.
	UserPersistence struct {
		Connection *gorm.DB
	}
)

// NewUserPersistence creates a new user persistence.
func NewUserPersistence() repository.UserRepository {
	return UserPersistence{
		Connection: getDBConnection(),
	}
}

// Create create a user.
func (userPersistence UserPersistence) Create(user model.User) (model.User, error) {
	result := userPersistence.Connection.
		Create(&user)

	return user, result.Error
}

func (userPersistence UserPersistence) GetAll() ([]model.User, error) {
	users := []model.User{}

	result := userPersistence.Connection.
		Find(&users)

	return users, result.Error
}

func (userPersistence UserPersistence) FindByID(id uint64) (model.User, error) {
	user := model.User{}

	result := userPersistence.Connection.
		Where(`"id" = ?`, id).
		Find(&user)

	return user, result.Error
}
