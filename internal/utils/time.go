package utils

import (
	"fmt"
	"time"
)

func ParseTime(s string) (time.Time, error) {
	layout := "2006-01-02 15:04:05.999999999 -0700"
	t, err := time.Parse(layout, s)
	if err != nil {
		return time.Time{}, fmt.Errorf("failed to parse time: %w", err)
	}
	return t, nil
}

func FormatDate(t time.Time) string {
	year, month, day := t.Date()
	hour, min, sec := t.Clock()
	millis := t.Nanosecond() / 1e6                  // Extract milliseconds
	fractional := fmt.Sprintf("%03d000000", millis) // Pad to 9 digits

	// Calculate timezone offset
	_, offsetSeconds := t.Zone()
	jsOffsetMinutes := -offsetSeconds / 60
	sign := "+"
	if jsOffsetMinutes > 0 {
		sign = "-"
	}
	absOffset := jsOffsetMinutes
	if absOffset < 0 {
		absOffset = -absOffset
	}
	offsetHours := absOffset / 60
	offsetMinutes := absOffset % 60

	// Format all parts with leading zeros
	monthStr := fmt.Sprintf("%02d", month)
	dayStr := fmt.Sprintf("%02d", day)
	hourStr := fmt.Sprintf("%02d", hour)
	minStr := fmt.Sprintf("%02d", min)
	secStr := fmt.Sprintf("%02d", sec)
	offsetHoursStr := fmt.Sprintf("%02d", offsetHours)
	offsetMinutesStr := fmt.Sprintf("%02d", offsetMinutes)

	return fmt.Sprintf("%d-%s-%s %s:%s:%s.%s %s%s%s",
		year, monthStr, dayStr, hourStr, minStr, secStr, fractional,
		sign, offsetHoursStr, offsetMinutesStr)
}
