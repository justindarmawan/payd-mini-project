package utils

import (
	"strings"
	"time"

	"github.com/justindarmawan/payd-mini-project/backend/viewmodel"
)

func validateShiftReq(input viewmodel.ShiftReq) []string {
	var errs []string

	if strings.TrimSpace(input.Date) != "" {
		shiftDate, err := time.Parse("2006-01-02", input.Date)
		if err != nil {
			errs = append(errs, "invalid date format (must be YYYY-MM-DD)")
		} else {
			today := time.Now().UTC().Truncate(24 * time.Hour)
			if shiftDate.Before(today) {
				errs = append(errs, "shift date must be today or later")
			}
		}
	}
	if strings.TrimSpace(input.StartTime) != "" {
		_, err := time.Parse("15:04", input.StartTime)
		if err != nil {
			errs = append(errs, "invalid start_time format (must be HH:mm)")
		}
	}
	if strings.TrimSpace(input.EndTime) != "" {
		_, err := time.Parse("15:04", input.EndTime)
		if err != nil {
			errs = append(errs, "invalid end_time format (must be HH:mm)")
		}
	}
	if strings.TrimSpace(input.StartTime) != "" && strings.TrimSpace(input.EndTime) != "" {
		startTime, err1 := time.Parse("15:04", input.StartTime)
		endTime, err2 := time.Parse("15:04", input.EndTime)
		if err1 == nil && err2 == nil {
			if !endTime.After(startTime) {
				errs = append(errs, "end_time must be after start_time")
			}
		}
	}

	return errs
}

func ValidateCreateShiftReq(input viewmodel.ShiftReq) []string {
	var errs []string
	if strings.TrimSpace(input.Date) == "" {
		errs = append(errs, "date is required")
	}
	if strings.TrimSpace(input.StartTime) == "" {
		errs = append(errs, "start_time is required")
	}
	if strings.TrimSpace(input.EndTime) == "" {
		errs = append(errs, "end_time is required")
	}
	if strings.TrimSpace(input.Role) == "" {
		errs = append(errs, "role is required")
	}
	errs = append(errs, validateShiftReq(input)...)
	return errs
}

func ValidateEditShiftReq(input viewmodel.ShiftReq) []string {
	var errs []string
	if strings.TrimSpace(input.Date) == "" &&
		strings.TrimSpace(input.StartTime) == "" &&
		strings.TrimSpace(input.EndTime) == "" &&
		strings.TrimSpace(input.Role) == "" &&
		strings.TrimSpace(input.Location) == "" {
		errs = append(errs, "at least one field must be provided for update")
		return errs
	}
	errs = append(errs, validateShiftReq(input)...)
	return errs
}
