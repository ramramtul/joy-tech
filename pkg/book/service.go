package book

import (
	"encoding/json"
	"io/ioutil"
	"joy-tech/models"
	"net/http"
)

type BookService struct{}

func GetBookList(url string, limit int, offset int) (*models.BookPagination, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	subjects := models.Subject{}

	jsonErr := json.Unmarshal(body, &subjects)
	if jsonErr != nil {
		return nil, err
	}

	models.BookList = subjects.Works

	pagination := &models.BookPagination{
		Data:   subjects.Works,
		Limit:  limit,
		Offset: offset,
	}

	return pagination, nil
}
