package handler_test

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/T-V-N/gourlshortener/internal/app"
	"github.com/T-V-N/gourlshortener/internal/handler"
	"github.com/T-V-N/gourlshortener/internal/storage"
	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
)

func Test_HandlerPostURL(t *testing.T) {
	type want struct {
		statusCode int
		response string
	}
	tests := []struct {
		name  string
		body []byte
		want want
	}{
		{
			name:    "regular link sent",
			body: []byte("https://youtube.com"),
			want: want{
				statusCode: http.StatusCreated,
				response:   "http://example.com/e62e2446",
			},
		},
		{
			name:    "Wrong URL passed",
			body: []byte(""),
			want: want{
				statusCode: http.StatusBadRequest,
				response:   "Wrong URL passed\n",
			},
		},
	}
	st := storage.NewStorage(map[string]string{})
	app := app.InitApp(st)
	hn := handler.InitHandler(app)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(tt.body))
			request.Header.Set("Host", "localhost:8080")
            w := httptest.NewRecorder()
            h := http.HandlerFunc(hn.HandlePostURL)
            h.ServeHTTP(w, request)
            res := w.Result()
            defer res.Body.Close()
            resBody, _ := io.ReadAll(res.Body)

			assert.Equal(t,  tt.want.response, string(resBody))
			assert.Equal(t,  tt.want.statusCode, res.StatusCode)

		})
	}
}
func Test_HandlerGetURL(t *testing.T) {
	type want struct {
		statusCode int
		location string
	}
	tests := []struct {
		name  string
		param string
		want want
	}{
		{
			name:    "get good link",
			param: "e62e2446",
			want: want{
				statusCode: http.StatusTemporaryRedirect,
				location:   "https://youtube.com",
			},
		},
		{
			name:    "no link",
			param: "",
			want: want{
				statusCode: http.StatusBadRequest,
				location:   "",
			},
		},
	}

	st := storage.NewStorage(map[string]string{"e62e2446":"https://youtube.com"})
	app := app.InitApp(st)
	hn := handler.InitHandler(app)


	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := httptest.NewRequest(http.MethodGet, "/"+tt.param, nil)
            w := httptest.NewRecorder()
			r := chi.NewRouter()
			r.Get("/",hn.HandleGetURL)
			r.Get("/{urlHash}",hn.HandleGetURL)
            r.ServeHTTP(w, request)
            res := w.Result()

			assert.Equal(t,  tt.want.location, res.Header.Get("Location"))
			assert.Equal(t,  tt.want.statusCode, res.StatusCode)

		})
	}
}
