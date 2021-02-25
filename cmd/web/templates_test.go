package main

import (
	"testing"
	"time"
)

func TestHumanDate(t *testing.T) {
	t.Parallel()
	//Gives the ability to test numerous instances of date
	testsHd := []struct {
		name     string
		time     time.Time
		expected string
	}{
		{
			name:     "UTC",
			time:     time.Date(2021, 02, 23, 0, 0, 0, 0, time.UTC),
			expected: "02.23.2021",
		},
		{
			name:     "Empty",
			time:     time.Time{},
			expected: "",
		},
		{
			name:     "CET",
			time:     time.Date(2021, 02, 23, 0, 0, 0, 0, time.FixedZone("CET", 1*60*60)),
			expected: "02.23.2021",
		},
	}

	for _, tt := range testsHd {
		t.Run(tt.name, func(t *testing.T) { //Runs each sub-test
			hD := humanDate(tt.time)

			if hD != tt.expected {
				t.Errorf("Expected: %q Output: %q", tt.expected, hD)
			}

		})
	}

}
