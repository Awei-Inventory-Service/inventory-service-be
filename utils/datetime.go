package utils

import "time"

// StartOfDay returns the start of the day (00:00:00) for the given time
func StartOfDay(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

// EndOfDay returns the end of the day (23:59:59) for the given time
func EndOfDay(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 999999999, t.Location())
}

func IsSameDay(first string, second string) bool {
	firstDate, err := time.Parse("2006-01-02", first)
	if err != nil {
		return false
	}

	secondDate, err := time.Parse("2006-01-02", second)
	if err != nil {
		return false
	}

	return firstDate.Year() == secondDate.Year() &&
		firstDate.Month() == secondDate.Month() &&
		firstDate.Day() == secondDate.Day()
}
