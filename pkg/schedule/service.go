package schedule

import (
	"errors"
	"joy-tech/models"
	"time"
)

func SchedulePickUpBook(request models.ScheduleRequest) (*models.Schedule, error) {
	// Check if book exist and available on Book Data
	var choosenBook *models.Book
	for _, v := range models.BookList {
		if v.Availability.OpenlibraryWork == request.EditionNumber {
			choosenBook = &v
			break
		}
	}

	if choosenBook == nil {
		return nil, errors.New("Book not found.")
	}

	if !choosenBook.Availability.AvailableToBorrow {
		return nil, errors.New("The book cannot be borrowed.")
	}

	if request.PickUpSchedule.Before(time.Now().Add(time.Minute * 60)) {
		return nil, errors.New("Pick up time must be at least one hour from now")
	}

	schedule := models.Schedule{
		EditionNumber:  request.EditionNumber,
		UserID:         request.UserID,
		PickUpSchedule: request.PickUpSchedule,
	}

	models.ScheduleData = append(models.ScheduleData, schedule)

	return &schedule, nil
}
