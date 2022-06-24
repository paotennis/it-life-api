package persistence

import (
	"gorm.io/gorm"

	"github.com/rikuhatano09/it-life-api/pkg/domain/model"
	"github.com/rikuhatano09/it-life-api/pkg/domain/repository"
)

type (
	// UserPersistence is a persistence to preserve users.
	EventPersistence struct {
		Connection *gorm.DB
	}
)

// NewUserPersistence creates a new user persistence.
func NewEventPersistence() repository.EventRepository {
	return EventPersistence{
		Connection: getDBConnection(),
	}
}

// Create create a user.
func (eventPersistence EventPersistence) Create(event model.Event) (model.Event, error) {
	result := eventPersistence.Connection.
		Create(&event)

	return event, result.Error
}
