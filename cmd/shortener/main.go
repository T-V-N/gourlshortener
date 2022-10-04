package main

import (
	"log"
	"net/http"

	"github.com/T-V-N/gourlshortener/internal/app"
	"github.com/T-V-N/gourlshortener/internal/config"
	"github.com/T-V-N/gourlshortener/internal/handler"
	"github.com/T-V-N/gourlshortener/internal/middleware/gzip"
	"github.com/T-V-N/gourlshortener/internal/storage"

	"github.com/go-chi/chi/v5"
)

func main() {
	cfg, err := config.Init()
	if err != nil {
		log.Panic("error: %w", err)
		return
	}

	st := storage.InitStorage(map[string]string{}, cfg)
	a := app.InitApp(st, cfg)
	h := handler.InitHandler(a)

	router := chi.NewRouter()
	router.Get("/{urlHash}", h.HandleGetURL)
	router.Post("/", h.HandlePostURL)
	router.Post("/api/shorten", h.HandleShortenURL)
	log.Panic(http.ListenAndServe(a.Config.ServerAddress, gzip.GzipHandle(router)))
}
