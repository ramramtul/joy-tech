package models

import "time"

var BookList []Book

type Book struct {
	Key              string           `json:"key"`
	Title            string           `json:"title"`
	FirstPublishYear int              `json:"first_publish_year"`
	Authors          []BookAuthor     `json:"authors"`
	Availability     BookAvailability `json:"availability"`
}

type BookAuthor struct {
	Key  string `json:"key"`
	Name string `json:"name"`
}

type BookAvailability struct {
	Status             string `json:"status"`
	AvailableToBorrow  bool   `json:"available_to_borrow"`
	ISBN               string `json:"isbn"`
	OpenlibraryWork    string `json:"openlibrary_work"`
	OpenlibraryEdition string `json:"openlibrary_edition"`
}

type Subject struct {
	Key         string    `json:"key"`
	Name        string    `json:"name"`
	SubjectType string    `json:"subject_type"`
	WorkCount   int       `json:"work_count"`
	Works       []Book    `json:"works"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type BookPagination struct {
	Data   []Book
	Limit  int
	Offset int
}
