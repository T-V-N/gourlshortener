package handler_test

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/T-V-N/gourlshortener/internal/app"
	"github.com/T-V-N/gourlshortener/internal/config"
	"github.com/T-V-N/gourlshortener/internal/handler"
	"github.com/T-V-N/gourlshortener/internal/storage"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
)

func Test_HandlerPostURL(t *testing.T) {
	type want struct {
		response   string
		statusCode int
	}

	tests := []struct {
		name string
		body []byte
		want want
	}{
		{
			name: "regular link sent",
			body: []byte("https://youtube.com"),
			want: want{
				statusCode: http.StatusCreated,
				response:   "http://example.com/e62e2446",
			},
		},
		{
			name: "Wrong URL passed",
			body: []byte(""),
			want: want{
				statusCode: http.StatusBadRequest,
				response:   "Wrong URL passed\n",
			},
		},
		{
			name: "Incorrect URL passed",
			body: []byte("ht_t_p://google.com"),
			want: want{
				statusCode: http.StatusBadRequest,
				response:   "Wrong URL passed\n",
			},
		},
	}
	cfg, _ := config.Init()
	st := storage.InitStorage(map[string]string{}, cfg)
	a := app.InitApp(st, cfg)
	hn := handler.InitHandler(a)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(tt.body))
			w := httptest.NewRecorder()
			h := http.HandlerFunc(hn.HandlePostURL)
			h.ServeHTTP(w, request)
			res := w.Result()
			defer res.Body.Close()
			resBody, _ := io.ReadAll(res.Body)

			assert.Equal(t, tt.want.response, string(resBody))
			assert.Equal(t, tt.want.statusCode, res.StatusCode)
		})
	}
}
func Test_HandlerGetURL(t *testing.T) {
	type want struct {
		location   string
		statusCode int
	}

	tests := []struct {
		name  string
		param string
		want  want
	}{
		{
			name:  "get good link",
			param: "e62e2446",
			want: want{
				statusCode: http.StatusTemporaryRedirect,
				location:   "https://youtube.com",
			},
		},
		{
			name:  "no link",
			param: "",
			want: want{
				statusCode: http.StatusBadRequest,
				location:   "",
			},
		},
	}

	cfg, _ := config.Init()
	st := storage.InitStorage(map[string]string{"e62e2446": "https://youtube.com"}, cfg)
	a := app.InitApp(st, cfg)
	hn := handler.InitHandler(a)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := httptest.NewRequest(http.MethodGet, "/"+tt.param, nil)
			w := httptest.NewRecorder()
			h := http.HandlerFunc(hn.HandleGetURL)
			ctx := chi.NewRouteContext()
			ctx.URLParams.Add("urlHash", tt.param)
			rctx := context.WithValue(request.Context(), chi.RouteCtxKey, ctx)
			request = request.WithContext(rctx)
			h.ServeHTTP(w, request)

			res := w.Result()
			res.Body.Close()

			assert.Equal(t, tt.want.location, res.Header.Get("Location"))
			assert.Equal(t, tt.want.statusCode, res.StatusCode)
		})
	}
}

func Test_HandlerShortenURL(t *testing.T) {
	type want struct {
		response   string
		statusCode int
	}

	tests := []struct {
		name string
		body []byte
		want want
	}{
		{
			name: "regular link sent",
			body: []byte("https://youtube.com"),
			want: want{
				statusCode: http.StatusCreated,
				response:   "http://example.com/e62e2446",
			},
		},
		{
			name: "Wrong URL passed",
			body: []byte(""),
			want: want{
				statusCode: http.StatusBadRequest,
				response:   "Wrong URL passed\n",
			},
		},
		{
			name: "Incorrect URL passed",
			body: []byte("ht_t_p://google.com"),
			want: want{
				statusCode: http.StatusBadRequest,
				response:   "Wrong URL passed\n",
			},
		},
	}

	cfg, _ := config.Init()
	st := storage.InitStorage(map[string]string{}, cfg)
	a := app.InitApp(st, cfg)
	hn := handler.InitHandler(a)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(tt.body))
			w := httptest.NewRecorder()
			h := http.HandlerFunc(hn.HandlePostURL)
			h.ServeHTTP(w, request)
			res := w.Result()
			defer res.Body.Close()
			resBody, _ := io.ReadAll(res.Body)

			assert.Equal(t, tt.want.response, string(resBody))
			assert.Equal(t, tt.want.statusCode, res.StatusCode)
		})
	}
}
