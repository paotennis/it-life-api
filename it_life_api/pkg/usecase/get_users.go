package usecase

import (
	"github.com/rikuhatano09/it-life-api/internal/infrastructure/persistence"
	"github.com/rikuhatano09/it-life-api/pkg/domain/contract"
)

func GetUsers() ([]contract.UserWithWeekDataResponse, error) {
	userPersistence := persistence.NewUserPersistence()
	eventPersistence := persistence.NewEventPersistence()
	response := []contract.UserWithWeekDataResponse{}

	users, err := userPersistence.GetAll()
	if err != nil {
		return []contract.UserWithWeekDataResponse{}, err
	}

	for _, user := range users {
		weekItems, err := eventPersistence.GetWeekItemByUserID(user.ID)
		if err != nil {
			return []contract.UserWithWeekDataResponse{}, err
		}

		response = append(response, contract.UserWithWeekDataResponse{
			UserID:   user.ID,
			Age:      user.Age,
			Role:     user.Role,
			Company:  user.Company,
			WeekData: weekItems,
		})
	}
	return response, nil
}
