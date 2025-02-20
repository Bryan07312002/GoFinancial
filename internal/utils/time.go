package utils

import (
	"errors"
	"fmt"
	"regexp"
	"time"
)

func ParseTimestamp(dateStr string) (time.Time, float64, error) {
	// Regular expression to match the timestamp and extract the duration
	re := regexp.MustCompile(`^([0-9\- :\.]+) (-?\d{4}) (-\d{2}) m=\+([0-9\.]+)$`)
	matches := re.FindStringSubmatch(dateStr)
	if matches != nil {
		// If the string matches the "m=+" format, process it separately
		timestampStr := matches[1] + " " + matches[2] // Extract the timestamp part
		durationStr := matches[4]                     // Extract the duration part

		// Define the layout for the timestamp part
		layout := "2006-01-02 15:04:05.999999999 -0700"
		parsedTime, err := time.Parse(layout, timestampStr)
		if err != nil {
			// Return a formatted error using errors.New
			return time.Time{}, 0, errors.New(fmt.Sprintf("error parsing timestamp: %v", err))
		}

		// Convert the duration part to a float64
		var duration float64
		_, err = fmt.Sscanf(durationStr, "%f", &duration)
		if err != nil {
			// Return a formatted error if duration parsing fails
			return time.Time{}, 0, errors.New(fmt.Sprintf("error parsing duration: %v", err))
		}

		return parsedTime, duration, nil
	}

	// Return a custom error if the format doesn't match
	return time.Time{}, 0, errors.New("unexpected format or extra text in string")
}
