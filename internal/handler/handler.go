package handler

import (
	"io"
	"net/http"
	"strings"

	"github.com/T-V-N/gourlshortener/internal/app"
)

type Handler struct {
	app *app.App
}

type URL struct {
    URL string `json:"URL"`
}

func InitHandler(app *app.App) *Handler {
	return &Handler{app}
}
func (h *Handler) HandleGetURL(w http.ResponseWriter, r *http.Request) {
	id := strings.Split(r.URL.Path, "/")[1]
	if id == "" {
		http.Error(w, "no short URL provided", http.StatusBadRequest)
		return
	}

	url, err := h.app.GetUrl(id) 
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Add("Location", url)
	w.WriteHeader(http.StatusTemporaryRedirect) 
}

func (h *Handler) HandlePostURL(w http.ResponseWriter, r *http.Request) { 
	body, err := io.ReadAll(r.Body)
	if err != nil  {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	hash, err := h.app.SaveURL(string(body))
	if err != nil {
		http.Error(w, "Wrong URL passed", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("http://"+r.Host+"/"+hash))
}
