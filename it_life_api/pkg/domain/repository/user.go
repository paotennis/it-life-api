package repository

import "github.com/rikuhatano09/it-life-api/pkg/domain/model"

type (
	UserRepository interface {
		Create(model.User) (model.User, error)
	}
)
