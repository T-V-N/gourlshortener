package handler

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/T-V-N/gourlshortener/internal/app"
	"github.com/T-V-N/gourlshortener/internal/storage"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
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
	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()

	id := chi.URLParam(r, "urlHash")

	if id == "" {
		http.Error(w, "no short URL provided", http.StatusBadRequest)
		return
	}

	url, err := h.app.GetURL(id, ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Add("Location", url)
	w.WriteHeader(http.StatusTemporaryRedirect)
}

func (h *Handler) HandlePostURL(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()

	body, err := io.ReadAll(r.Body)

	uid := r.Context().Value("uid")

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	hash, err := h.app.SaveURL(string(body), uid.(string), ctx)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == pgerrcode.UniqueViolation {
				w.WriteHeader(http.StatusConflict)
				_, err = w.Write([]byte(h.app.Config.BaseURL + "/" + hash))

				if err != nil {
					http.Error(w, "Unknown error", http.StatusInternalServerError)
					return
				}

				return
			}
		}

		http.Error(w, "Wrong URL passed", http.StatusBadRequest)

		return
	}

	w.WriteHeader(http.StatusCreated)

	_, err = w.Write([]byte(hash))
	if err != nil {
		log.Println(err.Error())
	}
}

func (h *Handler) HandleShortenURL(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()

	obj := URL{}
	if err := json.NewDecoder(r.Body).Decode(&obj); err != nil {
		http.Error(w, "Error while parsing URL", http.StatusBadRequest)
		return
	}

	uid := r.Context().Value("uid")

	hash, err := h.app.SaveURL(obj.URL, uid.(string), ctx)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == pgerrcode.UniqueViolation {
				w.Header().Set("content-type", "application/json")
				w.WriteHeader(http.StatusConflict)

				shortenedURL := ShortenResult{Result: h.app.Config.BaseURL + "/" + hash}

				err = json.NewEncoder(w).Encode(shortenedURL)
				if err != nil {
					http.Error(w, "Unknown error", http.StatusInternalServerError)
					return
				}
				return
			}
		}

		http.Error(w, "Wrong URL passed", http.StatusBadRequest)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)

	shortenedURL := ShortenResult{Result: hash}

	err = json.NewEncoder(w).Encode(shortenedURL)
	if err != nil {
		http.Error(w, "Something went wrong", http.StatusBadRequest)
		return
	}
}

func (h *Handler) HandleShortenBatchURL(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()

	obj := []storage.BatchURL{}

	if err := json.NewDecoder(r.Body).Decode(&obj); err != nil {
		http.Error(w, "Error while parsing URL", http.StatusBadRequest)
		return
	}

	uid := r.Context().Value("uid")

	urls, err := h.app.BatchSaveURL(ctx, obj, uid.(string))
	if err != nil {
		http.Error(w, "Wrong URL passed", http.StatusBadRequest)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)

	err = json.NewEncoder(w).Encode(urls)
	if err != nil {
		http.Error(w, "Something went wrong", http.StatusBadRequest)
		return
	}
}

func (h *Handler) HandleListURL(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()

	uid := r.Context().Value("uid")

	url, err := h.app.GetURLByUID(uid.(string), ctx)
	if err != nil {
		http.Error(w, "Error getting URLs ;(", http.StatusBadRequest)
		return
	}

	if len(url) == 0 {
		http.Error(w, "No content", http.StatusNoContent)
		return
	}

	w.Header().Set("content-type", "application/json")

	if err != json.NewEncoder(w).Encode(url) {
		http.Error(w, "Something went wrong", http.StatusBadRequest)
		return
	}
}

func (h *Handler) HandlePing(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()

	err := h.app.PingStorage(ctx)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
