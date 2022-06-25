package usecase

import (
	"github.com/rikuhatano09/it-life-api/internal/infrastructure/persistence"
	"github.com/rikuhatano09/it-life-api/pkg/domain/contract"
	"github.com/rikuhatano09/it-life-api/pkg/domain/model"
)

func CreateEvents(requestBody contract.EventPostRequestBody) ([]model.Event, error) {
	eventPersistence := persistence.NewEventPersistence()
	events := []model.Event{}

	for _, eventItem := range requestBody.Events {
		event := model.Event{
			UserID:      requestBody.UserID,
			Name:        eventItem.Name,
			Description: eventItem.Description,
			StartsAt:    eventItem.GetStartsAt(requestBody.Date),
			EndsAt:      eventItem.GetEndsAt(requestBody.Date),
		}

		event, err := eventPersistence.Create(event)
		if err != nil {
			return []model.Event{}, err
		}
		events = append(events, event)
	}
	return events, nil
}
