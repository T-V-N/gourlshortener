package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/T-V-N/gourlshortener/internal/app"
	"github.com/T-V-N/gourlshortener/internal/config"
	"github.com/T-V-N/gourlshortener/internal/handler"
	"github.com/T-V-N/gourlshortener/internal/middleware/auth"
	"github.com/T-V-N/gourlshortener/internal/middleware/gzip"
	"golang.org/x/crypto/acme/autocert"

	"github.com/T-V-N/gourlshortener/internal/storage"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

var (
	buildVersion string
	buildDate    string
	buildCommit  string
)

func outputBuildInfo() {
	if buildVersion == "" {
		fmt.Printf("Build version: %v", "N/A")
	} else {
		fmt.Printf("Build version: %v", buildVersion)
	}

	if buildDate == "" {
		fmt.Printf("Build date: %v", "N/A")
	} else {
		fmt.Printf("Build date: %v", buildDate)
	}

	if buildCommit == "" {
		fmt.Printf("Build commit: %v \n", "N/A")
	} else {
		fmt.Printf("Build commit: %v \n", buildCommit)
	}
}

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	defer stop()

	cfg, err := config.Init()
	if err != nil {
		log.Panic("error: %w", err)
	}

	st := storage.InitStorage(map[string]storage.URL{}, cfg)
	a := app.NewApp(ctx, st, cfg)
	a.Init()
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
	router.Get("/api/internal/stats", h.HandleGetStats)

	outputBuildInfo()

	server := http.Server{
		Handler: router,
		Addr:    cfg.ServerAddress,
	}

	if cfg.EnableHTTPS {
		manager := &autocert.Manager{
			Cache:      autocert.DirCache("cache-dir"),
			Prompt:     autocert.AcceptTOS,
			HostPolicy: autocert.HostWhitelist("localhost"),
		}

		server.Addr = ":443"
		server.TLSConfig = manager.TLSConfig()

		go func() {
			err := server.ListenAndServeTLS("", "")

			if err != nil {
				panic(err)
			}
		}()
	} else {
		go func() {
			err := server.ListenAndServe()

			if err != nil {
				fmt.Println(err)
			}
		}()
	}

	<-ctx.Done()
	stop()

	shutdownCtx, stopShutdownCtx := context.WithTimeout(context.Background(), 10*time.Second)
	defer stopShutdownCtx()

	err = server.Shutdown(shutdownCtx)

	if err != nil {
		fmt.Print("Unable to shutdown server in 10 secs")
	}

	defer st.KillConn()
}
