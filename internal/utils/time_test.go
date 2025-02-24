package utils

import (
	"testing"
	"time"
)

func TestParseTime(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    time.Time
		wantErr bool
	}{
		{
			name:  "valid with nanoseconds",
			input: "2023-10-05 14:30:45.123456789 +0300",
			want:  time.Date(2023, time.October, 5, 14, 30, 45, 123456789, time.FixedZone("", 3*60*60)),
		},
		{
			name:  "valid with milliseconds",
			input: "2021-12-25 00:00:00.500 -0500",
			want:  time.Date(2021, time.December, 25, 0, 0, 0, 500000000, time.FixedZone("", -5*60*60)),
		},
		{
			name:    "invalid date",
			input:   "2006-13-02 15:04:05.000 -0700",
			wantErr: true,
		},
		{
			name:  "valid zero fractional",
			input: "2023-01-01 00:00:00.000000000 +0000",
			want:  time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseTime(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseTime() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && !got.Equal(tt.want) {
				t.Errorf("ParseTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
