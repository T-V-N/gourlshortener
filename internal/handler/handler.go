package handler

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/T-V-N/gourlshortener/internal/app"
	"github.com/go-chi/chi/v5"
)

type Handler struct {
	app *app.App
}

type URL struct {
	URL string `json:url`
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
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	hash, err := h.app.SaveURL(string(body))
	if err != nil {
		http.Error(w, "Wrong URL passed", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)

	_, err = w.Write([]byte("http://" + r.Host + "/" + hash))
	if err != nil {
		panic(err.Error())
	}
}

func (h *Handler) HandleShortenURL(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	obj := URL{}
	if err2 := json.Unmarshal(body, &obj); err2 != nil {
		http.Error(w, "Error while parsing URL", http.StatusBadRequest)
		return
	}

	hash, err := h.app.SaveURL(obj.URL)
	if err != nil {
		http.Error(w, "Wrong URL passed", http.StatusBadRequest)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)

	shortenedURL := URL{URL: "http://" + r.Host + "/" + hash}

	jsonResBody, err := json.Marshal(shortenedURL)
	if err != nil {
		http.Error(w, "Something went wrong", http.StatusBadRequest)
		return
	}

	_, err = w.Write(jsonResBody)
	if err != nil {
		panic(err.Error())
	}
}
