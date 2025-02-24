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
