package book

import (
	"encoding/json"
	"fmt"
	"joy-tech/helper"
	"net/http"
)

type BookHandler struct{}

func NewBookHandler() *BookHandler {
	return &BookHandler{}
}

func (h *BookHandler) HandleGetBookList(w http.ResponseWriter, r *http.Request) {
	publishedStartIn := -1
	publishedEndIn := -1
	lim := 10
	off := 1

	// move to ENV
	url := fmt.Sprintf("https://openlibrary.org/subjects/love.json?")

	ebooks := r.URL.Query().Get("ebooks")
	if ebooks != "" {
		res, err := helper.ParseToBool(ebooks)
		if err != nil {
			http.Error(w, "Error when parsing boolean", http.StatusUnprocessableEntity)
			return
		}

		url = fmt.Sprintf("%s&ebooks=%t", url, res)
	}

	publishedStart := r.URL.Query().Get("published_start")
	if publishedStart != "" {
		start, err := helper.ParseToInt(publishedStart)
		if err != nil {
			http.Error(w, "Published year must be number", http.StatusUnprocessableEntity)
			return
		}

		publishedStartIn = start
	}

	publishedEnd := r.URL.Query().Get("published_end")
	if publishedEnd != "" {
		end, err := helper.ParseToInt(publishedEnd)
		if err != nil {
			http.Error(w, "Published year must be number", http.StatusUnprocessableEntity)
			return
		}

		publishedEndIn = end
	}

	if publishedStartIn > -1 && publishedEndIn > -1 {
		url = fmt.Sprintf("%s&published_in=%s-%s", url, publishedStart, publishedEnd)
	} else if publishedStartIn > -1 {
		url = fmt.Sprintf("%s&published_in=%s", url, publishedStart)
	} else {
		url = fmt.Sprintf("%s&published_in=0-%s", url, publishedEnd)
	}

	limit := r.URL.Query().Get("limit")
	if limit != "" {
		res, err := helper.ParseToInt(limit)
		if err != nil {
			http.Error(w, "Limit must be number", http.StatusUnprocessableEntity)
			return
		}

		url = fmt.Sprintf("%s&limit=%d", url, res)
		lim = res
	} else {
		url = fmt.Sprintf("%s&limit=%d", url, lim)
	}

	offset := r.URL.Query().Get("offset")
	if offset != "" {
		res, err := helper.ParseToInt(offset)
		if err != nil {
			http.Error(w, "Offset must be number", http.StatusUnprocessableEntity)
			return
		}

		url = fmt.Sprintf("%s&offset=%d", url, res)
		off = res
	} else {
		url = fmt.Sprintf("%s&offset=%d", url, off)
	}

	bookList, err := GetBookList(url, lim, off)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(bookList)
}
