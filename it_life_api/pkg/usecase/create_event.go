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

func CreateEvent(requestBody contract.EventPostRequestBody) (model.Event, error) {
	eventPersistence := persistence.NewEventPersistence()

	event := model.Event{
		UserID:      requestBody.UserId,
		Name:        requestBody.Events.Name,
		Description: requestBody.Events.Description,
		StartsAt:    StrToTime(requestBody.Date, requestBody.Events.StartTime),
		EndsAt:      StrToTime(requestBody.Date, requestBody.Events.EndTime),
	}

	event, err := eventPersistence.Create(event)
	if err != nil {
		return model.Event{}, err
	}
	return event, nil
}
