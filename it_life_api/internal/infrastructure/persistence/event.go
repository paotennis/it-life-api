package persistence

import (
	"gorm.io/gorm"

	"github.com/rikuhatano09/it-life-api/pkg/domain/model"
	"github.com/rikuhatano09/it-life-api/pkg/domain/repository"
)

type (
	// EventPersistence is a persistence to preserve events.
	EventPersistence struct {
		Connection *gorm.DB
	}
)

// NewEventPersistence creates a new event persistence.
func NewEventPersistence() repository.EventRepository {
	return EventPersistence{
		Connection: getDBConnection(),
	}
}

// Create create a event.
func (eventPersistence EventPersistence) Create(event model.Event) (model.Event, error) {
	result := eventPersistence.Connection.
		Create(&event)

	return event, result.Error
}
