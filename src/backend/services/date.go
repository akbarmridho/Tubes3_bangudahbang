package services

import (
	"time"
)

func GetDay(input string) (string, error) {
	layout := "02/01/2006 15:04:05"
	date := input + " 00:00:00"
	t, err := time.Parse(layout, date)
	if err != nil {
		return "I'm sorry I can't get the day of the date, the date inputted is invalid. Make sure it's DD/MM/YYYY", nil
	}
	return t.Weekday().String(), nil
}
