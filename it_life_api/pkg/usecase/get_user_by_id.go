package usecase

import (
	"strings"

	"github.com/rikuhatano09/it-life-api/internal/infrastructure/persistence"
	"github.com/rikuhatano09/it-life-api/pkg/domain/contract"
)

func GetUserByID(id uint64, startDateQuery string) (contract.UserWithWeekEventResponse, error) {
	userPersistence := persistence.NewUserPersistence()
	eventPersistence := persistence.NewEventPersistence()
	response := contract.UserWithWeekEventResponse{}

	user, err := userPersistence.FindByID(id)
	if err != nil {
		return contract.UserWithWeekEventResponse{}, err
	}

	response = contract.UserWithWeekEventResponse{
		UserID:    user.ID,
		StartDate: strings.ReplaceAll(startDateQuery, "-", "/"),
		Age:       user.Age,
		Role:      user.Role,
		Company:   user.Company,
		Monday:    []contract.EventItem{},
		Tuesday:   []contract.EventItem{},
		Wednesday: []contract.EventItem{},
		Thursday:  []contract.EventItem{},
		Friday:    []contract.EventItem{},
		Saturday:  []contract.EventItem{},
		Sunday:    []contract.EventItem{},
	}

	beginningWeekDate := convertStringToTime(startDateQuery)
	events, err := eventPersistence.GetByUserID(id, beginningWeekDate)
	if err != nil {
		return contract.UserWithWeekEventResponse{}, err
	}

	for _, event := range events {
		eventItem := contract.EventItem{
			Name:        event.Name,
			StartTime:   suppressZeroHour(event.StartsAt.Format("15:04")),
			EndTime:     suppressZeroHour(event.EndsAt.Format("15:04")),
			Description: event.Description,
		}
		switch event.StartsAt.Weekday() {
		case 0:
			response.Sunday = append(response.Sunday, eventItem)
		case 1:
			response.Monday = append(response.Monday, eventItem)
		case 2:
			response.Tuesday = append(response.Tuesday, eventItem)
		case 3:
			response.Wednesday = append(response.Wednesday, eventItem)
		case 4:
			response.Thursday = append(response.Thursday, eventItem)
		case 5:
			response.Friday = append(response.Friday, eventItem)
		case 6:
			response.Saturday = append(response.Saturday, eventItem)
		}
	}
	return response, nil
}
