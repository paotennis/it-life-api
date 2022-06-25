package repository

import "github.com/rikuhatano09/it-life-api/pkg/domain/model"

type (
	EventRepository interface {
		Create(model.Event) (model.Event, error)
	}
)
