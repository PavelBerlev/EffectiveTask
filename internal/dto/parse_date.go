package dto

import (
	"errors"
	"strings"
	"time"
)

func ParseMonthYear(s string) (time.Time, error) {
	t, err := time.Parse("01-2006", s)
	if err != nil {
		return time.Time{}, err
	}
	return time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, time.UTC), nil
}

func ValidateDateRange(startDateStr string, endDateStr *string) (time.Time, *time.Time, error) {
	start, err := ParseMonthYear(startDateStr)
	if err != nil {
		return time.Time{}, nil, err
	}

	var end *time.Time
	if endDateStr != nil && strings.TrimSpace(*endDateStr) != "" {
		e, err := ParseMonthYear(*endDateStr)
		if err != nil {
			return time.Time{}, nil, err
		}

		if !e.After(start) {
			return time.Time{}, nil, errors.New("end date must be after start date")
		}

		end = &e
	}

	return start, end, nil
}
