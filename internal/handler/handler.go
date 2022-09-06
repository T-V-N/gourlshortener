package handler

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/T-V-N/gourlshortener/internal/storage"
)

type Handler struct {
	storage *storage.Storage
}

type URL struct {
    URL string `json:"URL"`
}

func InitHandler(st *storage.Storage) *Handler {
	return &Handler{storage: st}
}
func (h *Handler) HandleGetURL(w http.ResponseWriter, r *http.Request) {
	id := strings.Split(r.URL.Path, "/")[1]
	if id == "" {
		http.Error(w, "no short URL provided", 500)
		return
	}

	url, err := h.storage.GetUrl(id) 
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Add("Location", url)
	w.WriteHeader(http.StatusTemporaryRedirect) 
}

func (h *Handler) HandlePostURL(w http.ResponseWriter, r *http.Request) { 
	url, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	hash := h.storage.SaveUrl(strings.ToLower(string(url)))
	resp, err := json.Marshal(URL{r.Host+"/"+hash})
	w.WriteHeader(http.StatusCreated)
	w.Write(resp)
}
