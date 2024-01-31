package book

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"joy-tech/models"
	"log"
	"net/http"
)

type BookService struct{}

func GetBookList(req models.GetBookRequest) ([]models.Book, error) {
	url := fmt.Sprintf("https://openlibrary.org/subjects/%s.json?details=%t&ebooks=%t&published_in=%s&limit=%d&offset=%d", req.Subject, req.Details, req.Ebooks, req.PublishedIn, req.Limit, req.Offset)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Print(err.Error())
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	subjects := models.Subject{}

	jsonErr := json.Unmarshal(body, &subjects)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	books := subjects.Works

	return books, nil
}
