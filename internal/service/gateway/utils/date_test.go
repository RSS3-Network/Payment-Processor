package utils

import (
	"testing"
	"time"
)

func TestGetUTCDateWithNow(t *testing.T) {
	// Getting the UTC date using the function
	currentTime := time.Now()
	utcDate := GetUTCDate(currentTime)

	// Getting the current time and setting it to 00:00:00 for comparison

	expectedDate := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 0, 0, 0, 0, time.UTC)

	// Checking if the year, month, and day are correct
	if utcDate.Year() != expectedDate.Year() || utcDate.Month() != expectedDate.Month() || utcDate.Day() != expectedDate.Day() {
		t.Errorf("Expected date to be %v but got %v", expectedDate, utcDate)
	}

	// Checking if the time is set to 00:00:00
	if utcDate.Hour() != 0 || utcDate.Minute() != 0 || utcDate.Second() != 0 {
		t.Errorf("Expected time to be 00:00:00 but got %02d:%02d:%02d", utcDate.Hour(), utcDate.Minute(), utcDate.Second())
	}

	// Checking if the location is set to UTC
	if utcDate.Location().String() != "UTC" {
		t.Errorf("Expected location to be UTC but got %v", utcDate.Location())
	}
}

func TestGetUTCDateWithStaticDate(t *testing.T) {
	// Setting a static date and time
	staticTime := time.Date(2023, 10, 26, 14, 30, 5, 1, time.UTC)

	// Getting the UTC date using the function
	utcDate := GetUTCDate(staticTime)

	// Setting the expected date based on the static date
	expectedDate := time.Date(2023, 10, 26, 0, 0, 0, 0, time.UTC)

	// Comparing the actual and expected dates
	if !utcDate.Equal(expectedDate) {
		t.Errorf("Expected date to be %v but got %v", expectedDate, utcDate)
	}
}
