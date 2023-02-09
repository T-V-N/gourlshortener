package handler_test

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/T-V-N/gourlshortener/internal/app"
	"github.com/T-V-N/gourlshortener/internal/handler"
	"github.com/T-V-N/gourlshortener/internal/storage"
)

func Example() {
	link := "https://yandex.ru"

	fmt.Print(link)

	cfg, _ := InitTestConfig()
	st := storage.InitStorage(map[string]storage.URL{}, cfg)
	a := app.NewApp(st, cfg)
	a.Init()
	hn := handler.InitHandler(a)

	request := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer([]byte(link)))
	w := httptest.NewRecorder()
	hn.HandlePostURL(w, request)

	resBody := w.Body.Bytes()

	fmt.Print(resBody)
}
