package services

import (
	"time"
)

func GetDay(input string) (string, error) {
	layoutFormat := "2006-01-02 15:04:05"
	value := input + " 08:04:00"
	date, err := time.Parse(layoutFormat, value)

	if err != nil {
		return "I'm sorry I can't get the day of the date, the date inputted is invalid. Make sure it's DD/MM/YYYY", nil
	}
	return input + " is " + date.Weekday().String(), nil
}
