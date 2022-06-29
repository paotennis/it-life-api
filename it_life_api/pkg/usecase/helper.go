package usecase

import (
	"strings"
	"time"
)

func convertStringToTime(value string) time.Time {
	timeValue, err := time.Parse("2006-01-02", value)
	if err != nil {
		// TODO: Custom error handling.
		panic(err)
	}
	return timeValue
}

func suppressZeroHour(value string) string {
	if strings.HasPrefix(value, "0") {
		return strings.TrimPrefix(value, "0")
	}
	return value
}
