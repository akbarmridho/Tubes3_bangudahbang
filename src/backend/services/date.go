package services

import (
	"time"
)

func GetDay(input string) (string, error) {
	var date time.Time
	found := false

	layoutFormats := []string{
		"2006-01-02 15:04:05",
		"2006-1-02 15:04:05",
		"2006-01-2 15:04:05",
		"2006-1-2 15:04:05",
		"02-01-2006 15:04:05",
		"2-01-2006 15:04:05",
		"02-1-2006 15:04:05",
		"2-1-2006 15:04:05",
	}

	for i := 0; i < len(layoutFormats) && !found; i++ {
		value := input + " 08:04:00"
		parsed, err := time.Parse(layoutFormats[i], value)

		if err == nil {
			date = parsed
			found = true
		}
	}

	if !found {
		return "Sorry the date you entered is invalid. Make sure it's in YYYY-MM-DD or DD-MM-YYYY format :D", nil
	}

	return input + " is " + date.Weekday().String(), nil
}
