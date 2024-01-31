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

func GetBookList(url string, limit int, offset int) (models.BookPagination, error) {
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

	models.BookList = subjects.Works

	pagination := models.BookPagination{
		Data:   subjects.Works,
		Limit:  limit,
		Offset: offset,
	}

	fmt.Print(models.BookList)

	return pagination, nil
}
