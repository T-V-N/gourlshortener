package main

import (
	"log"
	"net/http"

	"github.com/T-V-N/gourlshortener/internal/app"
	"github.com/T-V-N/gourlshortener/internal/handler"
	"github.com/T-V-N/gourlshortener/internal/storage"
	"github.com/go-chi/chi/v5"
)

func main() {
	st := storage.NewStorage(map[string]string{})
	a := app.InitApp(st)
	h := handler.InitHandler(a)

	router := chi.NewRouter()
	router.Get("/{urlHash}", h.HandleGetURL)
	router.Post("/", h.HandlePostURL)
	router.Post("/api/shorten", h.HandleShortenURL)
	log.Panic(http.ListenAndServe(":8080", router))
}
