package handler

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/T-V-N/gourlshortener/internal/app"
	"github.com/go-chi/chi/v5"
)

type Handler struct {
	app *app.App
}

type URL struct {
	URL string `json:"url"`
}

type ShortenResult struct {
	Result string `json:"result"`
}

func InitHandler(a *app.App) *Handler {
	return &Handler{a}
}

func (h *Handler) HandleGetURL(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "urlHash")

	if id == "" {
		http.Error(w, "no short URL provided", http.StatusBadRequest)
		return
	}

	url, err := h.app.GetURL(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Add("Location", url)
	w.WriteHeader(http.StatusTemporaryRedirect)
}

func (h *Handler) HandlePostURL(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)

	uid := r.Header.Get("uid")

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	hash, err := h.app.SaveURL(string(body), uid)
	if err != nil {
		http.Error(w, "Wrong URL passed", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)

	_, err = w.Write([]byte(h.app.Config.BaseURL + "/" + hash))
	if err != nil {
		log.Println(err.Error())
	}
}

func (h *Handler) HandleShortenURL(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	obj := URL{}
	if err := json.NewDecoder(bytes.NewBuffer(body)).Decode(&obj); err != nil {
		http.Error(w, "Error while parsing URL", http.StatusBadRequest)
		return
	}

	uid := r.Header.Get("uid")

	hash, err := h.app.SaveURL(obj.URL, uid)
	if err != nil {
		http.Error(w, "Wrong URL passed", http.StatusBadRequest)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)

	shortenedURL := ShortenResult{Result: h.app.Config.BaseURL + "/" + hash}

	jsonResBody, err := json.Marshal(shortenedURL)
	if err != nil {
		http.Error(w, "Something went wrong", http.StatusBadRequest)
		return
	}

	_, err = w.Write(jsonResBody)
	if err != nil {
		log.Println(err.Error())
	}
}

func (h *Handler) HandleListURL(w http.ResponseWriter, r *http.Request) {
	uid := r.Header.Get("uid")

	url, err := h.app.GetURLByUID(uid)
	if err != nil {
		http.Error(w, "Error getting URLs ;(", http.StatusBadRequest)
		return
	}

	if len(url) == 0 {
		http.Error(w, "No content", http.StatusNoContent)
		return
	}

	w.Header().Set("content-type", "application/json")

	jsonResBody, err := json.Marshal(url)

	if err != nil {
		http.Error(w, "Something went wrong", http.StatusBadRequest)
		return
	}

	_, err = w.Write(jsonResBody)
	if err != nil {
		log.Println(err.Error())
	}
}
