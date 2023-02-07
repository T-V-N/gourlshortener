package app_test

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/T-V-N/gourlshortener/internal/app"
	"github.com/T-V-N/gourlshortener/internal/config"
	"github.com/T-V-N/gourlshortener/internal/storage"
	"github.com/caarlos0/env/v6"

	"github.com/stretchr/testify/assert"
)

func InitTestConfig() (*config.Config, error) {
	cfg := &config.Config{}
	err := env.Parse(cfg)

	if err != nil {
		return nil, fmt.Errorf("error: %w", err)
	}

	return cfg, nil
}

func Test_PostURL(t *testing.T) {
	type want struct {
		statusCode int
		hash       string
	}

	tests := []struct {
		name string
		uid  string
		url  string
		want want
	}{
		{
			name: "ok URL sent",
			url:  "https://youtube.com",
			uid:  "1111",
			want: want{
				statusCode: http.StatusCreated,
				hash:       "http://localhost:8080/e62e2446",
			},
		},
		{
			name: "Wrong URL passed",
			url:  "",
			uid:  "1111",
			want: want{
				statusCode: http.StatusBadRequest,
				hash:       "Wrong URL passed\n",
			},
		},
		{
			name: "Incorrect URL passed",
			url:  "ht_t_p://google.com",
			uid:  "1111",
			want: want{
				statusCode: http.StatusBadRequest,
				hash:       "Wrong URL passed\n",
			},
		},
	}
	cfg, _ := InitTestConfig()
	st := storage.InitStorage(map[string]storage.URL{}, cfg)
	a := app.InitApp(st, cfg)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hash, _ := a.SaveURL(context.Background(), tt.url, tt.uid)

			assert.Equal(t, tt.want.hash, hash)
		})
	}
}

// func Test_GetURL(t *testing.T) {
// 	type want struct {
// 		location   string
// 		statusCode int
// 	}

// 	tests := []struct {
// 		name  string
// 		param string
// 		want  want
// 	}{
// 		{
// 			name:  "get good link",
// 			param: "e62e2446",
// 			want: want{
// 				statusCode: http.StatusTemporaryRedirect,
// 				location:   "https://youtube.com",
// 			},
// 		},
// 		{
// 			name:  "no link",
// 			param: "",
// 			want: want{
// 				statusCode: http.StatusBadRequest,
// 				location:   "",
// 			},
// 		},
// 	}

// 	cfg, _ := InitTestConfig()
// 	st := storage.InitStorage(map[string]storage.URL{"e62e2446": {UID: "", ShortURL: "e62e2446", URL: "https://youtube.com"}}, cfg)
// 	a := app.InitApp(st, cfg)
// 	hn := handler.InitHandler(a)

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			request := httptest.NewRequest(http.MethodGet, "/"+tt.param, nil)
// 			w := httptest.NewRecorder()
// 			ctx := chi.NewRouteContext()
// 			ctx.URLParams.Add("urlHash", tt.param)
// 			rctx := context.WithValue(request.Context(), chi.RouteCtxKey, ctx)
// 			request = request.WithContext(rctx)
// 			hn.HandleGetURL(w, request)

// 			res := w.Result()
// 			res.Body.Close()

// 			assert.Equal(t, tt.want.location, res.Header.Get("Location"))
// 			assert.Equal(t, tt.want.statusCode, res.StatusCode)
// 		})
// 	}
// }

// func Test_HandlerShortenURL(t *testing.T) {
// 	type want struct {
// 		response   string
// 		statusCode int
// 	}

// 	tests := []struct {
// 		name string
// 		body []byte
// 		want want
// 	}{
// 		{
// 			name: "regular link sent",
// 			body: []byte("https://youtube.com"),
// 			want: want{
// 				statusCode: http.StatusCreated,
// 				response:   "http://localhost:8080/e62e2446",
// 			},
// 		},
// 		{
// 			name: "Wrong URL passed",
// 			body: []byte(""),
// 			want: want{
// 				statusCode: http.StatusBadRequest,
// 				response:   "Wrong URL passed\n",
// 			},
// 		},
// 		{
// 			name: "Incorrect URL passed",
// 			body: []byte("ht_t_p://google.com"),
// 			want: want{
// 				statusCode: http.StatusBadRequest,
// 				response:   "Wrong URL passed\n",
// 			},
// 		},
// 	}

// 	cfg, _ := InitTestConfig()
// 	st := storage.InitStorage(map[string]storage.URL{}, cfg)
// 	app := app.InitApp(st, cfg)
// 	hn := handler.InitHandler(app)

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			request := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(tt.body))
// 			w := httptest.NewRecorder()
// 			hn.HandlePostURL(w, request)
// 			resBody := w.Body.Bytes()

// 			assert.Equal(t, tt.want.response, string(resBody))
// 			assert.Equal(t, tt.want.statusCode, w.Code)
// 		})
// 	}
// }

// func Test_HandleShortenBatchURL(t *testing.T) {
// 	type want struct {
// 		response   []storage.BatchURL
// 		statusCode int
// 	}

// 	tests := []struct {
// 		name string
// 		body []storage.BatchURL
// 		want want
// 	}{
// 		{
// 			name: "regular link sent",
// 			body: []storage.BatchURL{
// 				{OriginalURL: "http://yandex.ru", CorrelationID: "js21y3", ShortURL: ""},
// 				{OriginalURL: "http://google.com", CorrelationID: "zxfjasd", ShortURL: ""},
// 			},
// 			want: want{
// 				statusCode: http.StatusCreated,
// 				response: []storage.BatchURL{
// 					{OriginalURL: "http://yandex.ru", CorrelationID: "js21y3", ShortURL: "http://localhost:8080/js21y3"},
// 					{OriginalURL: "http://google.com", CorrelationID: "js21y3", ShortURL: "http://localhost:8080/zxfjasd"},
// 				},
// 			},
// 		},
// 	}

// 	cfg, _ := InitTestConfig()
// 	st := storage.InitStorage(map[string]storage.URL{}, cfg)
// 	app := app.InitApp(st, cfg)
// 	hn := handler.InitHandler(app)

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			body := bytes.NewBuffer([]byte{})
// 			json.NewEncoder(body).Encode(tt.body)
// 			request := httptest.NewRequest(http.MethodPost, "/api/shorten/batch", body)

// 			w := httptest.NewRecorder()

// 			hn.HandleShortenBatchURL(w, request)

// 			resp := []storage.BatchURL{}
// 			json.NewDecoder(w.Body).Decode(&resp)

// 			for i, el := range resp {
// 				assert.Equal(t, tt.want.response[i].ShortURL, el.ShortURL)
// 			}
// 		})
// 	}
// }
