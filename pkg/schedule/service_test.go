package schedule

import (
	"joy-tech/models"
	"testing"
	"time"
)

func TestSchedulePickUpBook(t *testing.T) {
	// Mock data for testing
	models.BookList = []models.Book{
		{Availability: models.BookAvailability{OpenlibraryWork: "OL10278W", AvailableToBorrow: true}},
		{Availability: models.BookAvailability{OpenlibraryWork: "OL12345X", AvailableToBorrow: false}},
	}

	tests := []struct {
		name           string
		request        models.ScheduleRequest
		expectedResult *models.Schedule
		expectedError  string
	}{
		{
			name: "Valid Request",
			request: models.ScheduleRequest{
				EditionNumber:  "OL10278W",
				UserID:         1,
				PickUpSchedule: time.Now().Add(time.Hour * 2), // Valid pick-up time (2 hours from now)
			},
			expectedResult: &models.Schedule{
				EditionNumber:  "OL10278W",
				UserID:         1,
				PickUpSchedule: time.Now().Add(time.Hour * 2),
			},
			expectedError: "",
		},
		{
			name: "Book Not Found",
			request: models.ScheduleRequest{
				EditionNumber:  "NonExistentEdition",
				UserID:         1,
				PickUpSchedule: time.Now().Add(time.Hour * 2),
			},
			expectedResult: nil,
			expectedError:  "Book not found.",
		},
		{
			name: "Book Not Available",
			request: models.ScheduleRequest{
				EditionNumber:  "OL12345X",
				UserID:         1,
				PickUpSchedule: time.Now().Add(time.Hour * 2),
			},
			expectedResult: nil,
			expectedError:  "The book cannot be borrowed.",
		},
		{
			name: "Invalid Pick-up Time",
			request: models.ScheduleRequest{
				EditionNumber:  "OL10278W",
				UserID:         1,
				PickUpSchedule: time.Now().Add(-time.Minute), // Invalid pick-up time (1 minute ago)
			},
			expectedResult: nil,
			expectedError:  "Pick up time must be at least one hour from now",
		},
		// Add more test cases as needed
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := SchedulePickUpBook(tt.request)

			// Check the result
			if tt.expectedResult != nil {
				if result == nil {
					t.Error("Expected a non-nil result, but got nil")
				}
			} else {
				if result != nil {
					t.Error("Expected nil result, but got non-nil")
				}
			}

			// Check the error
			if tt.expectedError != "" {
				if err == nil {
					t.Error("Expected an error, but got nil")
				} else if err.Error() != tt.expectedError {
					t.Errorf("Expected error: %s, but got: %s", tt.expectedError, err.Error())
				}
			} else {
				if err != nil {
					t.Errorf("Expected no error, but got: %s", err.Error())
				}
			}
		})
	}
}
