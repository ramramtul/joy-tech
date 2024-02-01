package schedule

import (
	"bytes"
	"joy-tech/pkg/book"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandleSchedulePickUp(t *testing.T) {
	tests := []struct {
		name            string
		requestBody     string
		expectedCode    int
		expectedError   string
		expectedPayload string // Optional: Expected response payload
	}{
		{
			name:            "Success",
			requestBody:     `{"edition_number": "OL10278W", "user_id": 1, "pick_up_schedule": "2025-01-31T10:48:25Z"}`,
			expectedCode:    http.StatusOK,
			expectedError:   "",
			expectedPayload: `{"edition_number":"OL10278W","user_id":1,"pick_up_schedule":"2025-01-31T10:48:25Z"}`,
		},
		{
			name:            "Invalid edition number",
			requestBody:     `{"edition_number": "invalid", "user_id": 1, "pick_up_schedule": "2024-01-31T10:48:25Z"}`,
			expectedCode:    http.StatusBadRequest,
			expectedError:   "",
			expectedPayload: `Book not found.`,
		},
		{
			name:            "Invalid JSON",
			requestBody:     `invalid-json`,
			expectedCode:    http.StatusBadRequest,
			expectedError:   "invalid character",
			expectedPayload: "",
		},
		{
			name:            "Empty Fields",
			requestBody:     `{"edition_number": "", "user_id": 0}`,
			expectedCode:    http.StatusUnprocessableEntity,
			expectedError:   "Fields cannot be empty",
			expectedPayload: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pp := httptest.NewRecorder()

			get, err := http.NewRequest("GET", "/book/list", nil)
			if err != nil {
				t.Fatalf("fail to create request: %s", err.Error())
			}

			q := get.URL.Query()
			q.Add("ebooks", "true")
			q.Add("published_start", "2000")
			q.Add("published_end", "2005")
			q.Add("limit", "10")
			q.Add("offset", "1")
			get.URL.RawQuery = q.Encode()

			book.HandleGetBookList(pp, get)

			// Create a request with the provided test case body
			req, err := http.NewRequest("POST", "/schedule/pickup", bytes.NewBufferString(tt.requestBody))
			if err != nil {
				t.Fatal(err)
			}

			// Create a ResponseRecorder to record the response
			rr := httptest.NewRecorder()

			// Call the handler function with the created request and response recorder
			HandleSchedulePickUp(rr, req)

			// Check the response status code
			if status := rr.Code; status != tt.expectedCode {
				t.Errorf("Unexpected status code: got %v want %v", status, tt.expectedCode)
			}

			// Check the response body
			if tt.expectedError != "" {
				responseBody := strings.TrimSpace(rr.Body.String())
				if !strings.Contains(responseBody, tt.expectedError) {
					t.Errorf("Expected error %q not found in response body: %q", tt.expectedError, responseBody)
				}
			}

			// Optionally, check the response payload
			if tt.expectedPayload != "" {
				expectedPayload := strings.TrimSpace(tt.expectedPayload)
				responseBody := strings.TrimSpace(rr.Body.String())
				if responseBody != expectedPayload {
					t.Errorf("Unexpected response payload:\nGot: %v\nWant: %v", responseBody, expectedPayload)
				}
			}
		})
	}
}
