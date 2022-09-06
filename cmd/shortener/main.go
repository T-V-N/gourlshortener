package main

import (
	"fmt"
	"net/http"

	"github.com/T-V-N/gourlshortener/internal/app"
	"github.com/T-V-N/gourlshortener/internal/handler"
	"github.com/T-V-N/gourlshortener/internal/storage"
	"github.com/go-chi/chi/v5"
)

func main() {
	st := storage.NewStorage(map[string]string{})
	app := app.InitApp(st)
	h := handler.InitHandler(app)
	
	router := chi.NewRouter()
	router.Get("/:id", h.HandleGetURL)
	router.Post("/", h.HandlePostURL)
	http.ListenAndServe(":8080", router)
	fmt.Println("go!")
}