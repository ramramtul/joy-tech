package book

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleGetBookList(t *testing.T) {
	type args struct {
		req *http.Request
	}

	invalidData := "invalid"

	tests := []struct {
		name     string
		args     func(t *testing.T) args
		wantCode int
		wantBody string
	}{
		{
			name: "must return http.StatusOK",
			args: func(*testing.T) args {
				req, err := http.NewRequest("GET", "/book/list", nil)
				if err != nil {
					t.Fatalf("fail to create request: %s", err.Error())
				}

				q := req.URL.Query()
				q.Add("ebooks", "true")
				q.Add("published_start", "2000")
				q.Add("published_end", "2005")
				q.Add("limit", "2")
				q.Add("offset", "1")
				req.URL.RawQuery = q.Encode()

				return args{
					req: req,
				}
			},
			wantCode: http.StatusOK,
		},
		{
			name: "must return http.StatusUnprocessableEntity cause invalid ebook param",
			args: func(*testing.T) args {
				req, err := http.NewRequest("GET", "/book/list", nil)
				if err != nil {
					t.Fatalf("fail to create request: %s", err.Error())
				}

				q := req.URL.Query()
				q.Add("ebooks", invalidData)
				req.URL.RawQuery = q.Encode()

				return args{
					req: req,
				}
			},
			wantCode: http.StatusUnprocessableEntity,
		},
		{
			name: "must return http.StatusUnprocessableEntity cause invalid offset param",
			args: func(*testing.T) args {
				req, err := http.NewRequest("GET", "/book/list", nil)
				if err != nil {
					t.Fatalf("fail to create request: %s", err.Error())
				}

				q := req.URL.Query()
				q.Add("ebooks", "true")
				q.Add("published_start", "2000")
				q.Add("published_end", "2005")
				q.Add("limit", "2")
				q.Add("offset", invalidData)
				req.URL.RawQuery = q.Encode()

				return args{
					req: req,
				}
			},
			wantCode: http.StatusUnprocessableEntity,
		},
		{
			name: "must return http.StatusUnprocessableEntity cause invalid limit param",
			args: func(*testing.T) args {
				req, err := http.NewRequest("GET", "/book/list", nil)
				if err != nil {
					t.Fatalf("fail to create request: %s", err.Error())
				}

				q := req.URL.Query()
				q.Add("ebooks", "true")
				q.Add("published_start", "2000")
				q.Add("published_end", "2005")
				q.Add("limit", invalidData)
				req.URL.RawQuery = q.Encode()

				return args{
					req: req,
				}
			},
			wantCode: http.StatusUnprocessableEntity,
		},
		{
			name: "must return http.StatusUnprocessableEntity cause invalid published_end param",
			args: func(*testing.T) args {
				req, err := http.NewRequest("GET", "/book/list", nil)
				if err != nil {
					t.Fatalf("fail to create request: %s", err.Error())
				}

				q := req.URL.Query()
				q.Add("ebooks", "true")
				q.Add("published_start", "2000")
				q.Add("published_end", invalidData)
				req.URL.RawQuery = q.Encode()

				return args{
					req: req,
				}
			},
			wantCode: http.StatusUnprocessableEntity,
		},
		{
			name: "must return http.StatusUnprocessableEntity cause invalid published_end param",
			args: func(*testing.T) args {
				req, err := http.NewRequest("GET", "/book/list", nil)
				if err != nil {
					t.Fatalf("fail to create request: %s", err.Error())
				}

				q := req.URL.Query()
				q.Add("ebooks", "true")
				q.Add("published_start", invalidData)
				req.URL.RawQuery = q.Encode()

				return args{
					req: req,
				}
			},
			wantCode: http.StatusUnprocessableEntity,
		},
		{
			name: "must return http.StatusOK without published_end",
			args: func(*testing.T) args {
				req, err := http.NewRequest("GET", "/book/list", nil)
				if err != nil {
					t.Fatalf("fail to create request: %s", err.Error())
				}

				q := req.URL.Query()
				q.Add("ebooks", "true")
				q.Add("published_start", "2000")
				q.Add("limit", "2")
				q.Add("offset", "1")
				req.URL.RawQuery = q.Encode()

				return args{
					req: req,
				}
			},
			wantCode: http.StatusOK,
		},
		{
			name: "must return http.StatusOK without published_start",
			args: func(*testing.T) args {
				req, err := http.NewRequest("GET", "/book/list", nil)
				if err != nil {
					t.Fatalf("fail to create request: %s", err.Error())
				}

				q := req.URL.Query()
				q.Add("ebooks", "true")
				q.Add("published_end", "2000")
				q.Add("limit", "2")
				q.Add("offset", "1")
				req.URL.RawQuery = q.Encode()

				return args{
					req: req,
				}
			},
			wantCode: http.StatusOK,
		},
		{
			name: "must return http.StatusOK without limit",
			args: func(*testing.T) args {
				req, err := http.NewRequest("GET", "/book/list", nil)
				if err != nil {
					t.Fatalf("fail to create request: %s", err.Error())
				}

				q := req.URL.Query()
				q.Add("ebooks", "true")
				q.Add("published_end", "2000")
				q.Add("offset", "1")
				req.URL.RawQuery = q.Encode()

				return args{
					req: req,
				}
			},
			wantCode: http.StatusOK,
		},
		{
			name: "must return http.StatusOK without offset",
			args: func(*testing.T) args {
				req, err := http.NewRequest("GET", "/book/list", nil)
				if err != nil {
					t.Fatalf("fail to create request: %s", err.Error())
				}

				q := req.URL.Query()
				q.Add("ebooks", "true")
				q.Add("published_end", "2000")
				req.URL.RawQuery = q.Encode()

				return args{
					req: req,
				}
			},
			wantCode: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tArgs := tt.args(t)
			resp := httptest.NewRecorder()
			handler := http.HandlerFunc(HandleGetBookList)
			handler.ServeHTTP(resp, tArgs.req)

			if resp.Result().StatusCode != tt.wantCode {
				t.Fatalf("the status code should be [%d] but received [%d]", resp.Result().StatusCode, tt.wantCode)
			}
		})
	}
}
