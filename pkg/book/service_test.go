package book

import (
	"joy-tech/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetBookList(t *testing.T) {
	testCases := []struct {
		name          string
		serverHandler http.HandlerFunc
		expectedData  []models.Book
		expectedError bool
		readAllError  error
		url           string
	}{
		{
			name: "Successful Response",
			serverHandler: func(w http.ResponseWriter, r *http.Request) {
			},
			expectedData:  []models.Book{{Title: "Book1"}, {Title: "Book2"}},
			expectedError: false,
			readAllError:  nil,
			url:           "https://openlibrary.org/subjects/love.json?&ebooks=true&published_in=2000-2005&limit=2&offset=1",
		},
		{
			name: "Error when url wrong",
			serverHandler: func(w http.ResponseWriter, r *http.Request) {
				w.Write([]byte(`{}`))
			},
			expectedData:  nil,
			url:           "www.example.com",
			readAllError:  nil,
			expectedError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			server := httptest.NewServer(tc.serverHandler)
			defer server.Close()

			url := server.URL
			if tc.url != "invalid" {
				url = tc.url
			}
			limit := 10
			offset := 0

			pagination, err := GetBookList(url, limit, offset)

			if tc.expectedError {
				assert.Error(t, err, "Expected an error")
				if tc.readAllError != nil {
					assert.Contains(t, err.Error(), tc.readAllError.Error(), "Expected specific error")
				}
			} else {
				assert.Nil(t, err, "Unexpected error")
				assert.NotNil(t, pagination, "Pagination should not be nil")
				assert.Equal(t, limit, pagination.Limit, "Limit should match")
				assert.Equal(t, offset, pagination.Offset, "Offset should match")
			}
		})
	}
}
