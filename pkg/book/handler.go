package book

import (
	"encoding/json"
	"joy-tech/helper"
	"joy-tech/models"
	"log"
	"net/http"
)

type BookHandler struct{}

func NewBookHandler() *BookHandler {
	return &BookHandler{}
}

func (h *BookHandler) HandleGetBookList(w http.ResponseWriter, r *http.Request) {
	var detBool bool
	var ebookBool bool
	var limInt int
	var offsetInt int

	subject := r.URL.Query().Get("subject")
	details := r.URL.Query().Get("details")
	if details != "" {
		res, err := helper.ParseToBool(details)
		if err != nil {
			log.Fatal(err)
		}

		detBool = res
	}

	ebooks := r.URL.Query().Get("ebooks")
	if ebooks != "" {
		res, err := helper.ParseToBool(ebooks)
		if err != nil {
			log.Fatal(err)
		}

		ebookBool = res
	}

	publishedIn := r.URL.Query().Get("publishedIn")
	limit := r.URL.Query().Get("limit")
	if limit != "" {
		res, err := helper.ParseToInt(limit)
		if err != nil {
			log.Fatal(err)
		}

		limInt = res
	}

	offset := r.URL.Query().Get("offset")
	if offset != "" {
		res, err := helper.ParseToInt(offset)
		if err != nil {
			log.Fatal(err)
		}

		offsetInt = res
	}

	req := models.GetBookRequest{
		Subject:     subject,
		Details:     detBool,
		Ebooks:      ebookBool,
		PublishedIn: publishedIn,
		Limit:       limInt,
		Offset:      offsetInt,
	}

	bookList, err := GetBookList(req)
	if err != nil {
		log.Fatal(err)
	}

	// Convert the user object to JSON and send it in the response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(bookList)

}
