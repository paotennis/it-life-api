package repository

import "github.com/rikuhatano09/it-life-api/pkg/domain/model"

type (
	UserRepository interface {
		Create(model.User) (model.User, error)
		GetAll() ([]model.User, error)
		FindByID(uint64) (model.User, error)
		FindByFirebaseUID(string) (model.User, error)
	}
)
