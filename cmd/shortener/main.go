package main

import (
	"log"
	"net/http"

	"github.com/T-V-N/gourlshortener/internal/app"
	"github.com/T-V-N/gourlshortener/internal/config"
	"github.com/T-V-N/gourlshortener/internal/handler"
	"github.com/T-V-N/gourlshortener/internal/middleware/auth"
	"github.com/T-V-N/gourlshortener/internal/middleware/gzip"

	"github.com/T-V-N/gourlshortener/internal/storage"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	cfg, err := config.Init()
	if err != nil {
		log.Panic("error: %w", err)
	}

	st := storage.InitStorage(map[string]storage.URL{}, cfg)
	a := app.InitApp(st, cfg)
	h := handler.InitHandler(a)
	authMw := auth.InitAuth(cfg)

	router := chi.NewRouter()

	router.Use(gzip.GzipHandle)
	router.Use(authMw)
	router.Use(middleware.Compress(5))
	router.Mount("/debug", middleware.Profiler())
	router.Get("/{urlHash}", h.HandleGetURL)
	router.Post("/", h.HandlePostURL)
	router.Post("/api/shorten", h.HandleShortenURL)
	router.Get("/api/user/urls", h.HandleListURL)
	router.Delete("/api/user/urls", h.HandleDeleteListURL)
	router.Post("/api/shorten/batch", h.HandleShortenBatchURL)
	router.Get("/ping", h.HandlePing)

	log.Panic(http.ListenAndServe(a.Config.ServerAddress, router))

	defer st.KillConn()
}
