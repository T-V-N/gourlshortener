package auth_test

import (
	"bytes"
	"encoding/json"

	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/T-V-N/gourlshortener/internal/app"
	"github.com/T-V-N/gourlshortener/internal/config"
	"github.com/T-V-N/gourlshortener/internal/handler"
	"github.com/T-V-N/gourlshortener/internal/middleware/auth"
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

var authH http.Handler
var rawCookie, respHash string

func Test_AuthHandler(t *testing.T) {
	cfg, _ := InitTestConfig()
	st := storage.InitStorage(map[string]storage.URL{}, cfg)
	a := app.InitApp(st, cfg)
	hn := handler.InitHandler(a)
	authH := auth.InitAuth(cfg)

	t.Run("request without cookie, should get a new cookie", func(t *testing.T) {
		nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			hn.HandlePostURL(w, r)
		})

		url := "https://yandexx.ru"

		request := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer([]byte(url)))

		w := httptest.NewRecorder()

		next := authH(nextHandler)
		next.ServeHTTP(w, request)

		rawCookie = w.Header().Get("Set-Cookie")
		respHash = w.Body.String()

		assert.NotEmpty(t, rawCookie)
		assert.Equal(t, http.StatusCreated, w.Code) //just to check if the underlying hanglers is ok
	})
	t.Run("request with auth cookie, should return the previously saved link", func(t *testing.T) {
		nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			hn.HandleListURL(w, r)
		})

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		request.Header.Set("cookie", rawCookie)

		w := httptest.NewRecorder()

		next := authH(nextHandler)
		next.ServeHTTP(w, request)

		resp := []storage.BatchURL{}
		json.NewDecoder(w.Body).Decode(&resp)

		assert.Equal(t, respHash, resp[0].ShortURL)
		assert.Equal(t, http.StatusOK, w.Code)
	})
	t.Run("request without auth cookie to url/list should return no content string", func(t *testing.T) {
		nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			hn.HandleListURL(w, r)
		})

		request := httptest.NewRequest(http.MethodGet, "/", nil)

		w := httptest.NewRecorder()

		next := authH(nextHandler)
		next.ServeHTTP(w, request)

		assert.Equal(t, w.Body.String(), "No content\n")
		assert.Equal(t, http.StatusNoContent, w.Code)
	})
}
