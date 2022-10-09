package gzip_test

import (
	"bytes"
	gzip "compress/gzip"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/T-V-N/gourlshortener/internal/app"
	"github.com/T-V-N/gourlshortener/internal/config"
	"github.com/T-V-N/gourlshortener/internal/handler"
	gzipMw "github.com/T-V-N/gourlshortener/internal/middleware/gzip"
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

func Test_GzipHandle(t *testing.T) {
	cfg, _ := InitTestConfig()
	st := storage.InitStorage(map[string]string{}, cfg)
	a := app.InitApp(st, cfg)
	hn := handler.InitHandler(a)

	t.Run("check gzipped request", func(t *testing.T) {
		nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			hn.HandlePostURL(w, r)
		})

		url := "https://yandexx.ru"

		buf := bytes.Buffer{}
		wr := gzip.NewWriter(&buf)
		wr.Write([]byte(url))
		wr.Close()

		request := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(buf.Bytes()))
		request.Header.Set("Accept-Encoding", "gzip")
		request.Header.Set("Content-Encoding", "gzip")

		w := httptest.NewRecorder()

		next := gzipMw.GzipHandle(nextHandler)
		next.ServeHTTP(w, request)

		resBody := w.Body.Bytes()

		assert.Equal(t, string(resBody), "http://localhost:8080/8f2eff5e")
		assert.Equal(t, http.StatusCreated, w.Code)
	})
	t.Run("check ungzipped request", func(t *testing.T) {
		nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			hn.HandlePostURL(w, r)
		})

		url := "https://google.com"

		request := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer([]byte(url)))

		w := httptest.NewRecorder()

		next := gzipMw.GzipHandle(nextHandler)
		next.ServeHTTP(w, request)

		resBody := w.Body.Bytes()

		assert.Equal(t, string(resBody), "http://localhost:8080/99999ebc")
		assert.Equal(t, http.StatusCreated, w.Code)
	})

}
