package repository

import (
	"time"

	"github.com/rikuhatano09/it-life-api/pkg/domain/contract"
	"github.com/rikuhatano09/it-life-api/pkg/domain/model"
)

type (
	EventRepository interface {
		Create(model.Event) (model.Event, error)
		GetWeekItemByUserID(uint64) ([]contract.WeekItem, error)
		GetByUserID(uint64, time.Time) ([]model.Event, error)
	}
)
