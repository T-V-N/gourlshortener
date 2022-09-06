package handler

import (
	"io"
	"net/http"
	"net/url"
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
		http.Error(w, "no short URL provided", http.StatusBadRequest)
		return
	}

	url, err := h.storage.GetUrl(id) 
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

	u, err := url.ParseRequestURI(string(body))
	if err != nil {
		http.Error(w, "Wrong URL passed", http.StatusBadRequest)
		return
	}
<<<<<<< HEAD
	hash := h.storage.SaveUrl(strings.ToLower(u.String()))
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("http://"+r.Host+"/"+hash))
=======
	hash := h.storage.SaveUrl(strings.ToLower(string(url)))
	
	w.WriteHeader(http.StatusCreated)
<<<<<<< HEAD
	w.Write([]byte(r.Host+"/"+hash))
>>>>>>> 79494da... remove json response body
=======
	w.Write([]byte("http://"+r.Host+"/"+hash))
>>>>>>> 007597d... proper localhost link
}
