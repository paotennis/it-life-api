package usecase

import (
	"time"

	"github.com/rikuhatano09/it-life-api/internal/infrastructure/persistence"
	"github.com/rikuhatano09/it-life-api/pkg/domain/contract"
	"github.com/rikuhatano09/it-life-api/pkg/domain/model"
)

func StrToTime(day string, times string) time.Time {
	str := day + times
	parsedTime, _ := time.Parse("2000-01-01T00:00:00Z", str)

	return parsedTime
}

func CreateEvent(requestBody contract.EventPostRequestBody) (contract.EventPostRequestBody, model.Event, error) {
	eventPersistence := persistence.NewEventPersistence()

	for i := 0; i < len(requestBody.Events); i++ {
		event := model.Event{
			UserID:      requestBody.UserId,
			Name:        requestBody.Events[i].Name,
			Description: requestBody.Events[i].Description,
			StartsAt:    StrToTime(requestBody.Date, requestBody.Events[i].StartTime),
			EndsAt:      StrToTime(requestBody.Date, requestBody.Events[i].EndTime),
		}

		event, err := eventPersistence.Create(event)
		if err != nil {
			return requestBody, model.Event{}, err
		}
	}
	return requestBody, model.Event{}, nil
}
