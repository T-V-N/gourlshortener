package gzip

import (
	"compress/gzip"
	"io"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5/middleware"
)

type Gzip struct {
	GzipReader gzip.Reader
	Body       io.ReadCloser
}

func (g *Gzip) Close() {
	g.GzipReader.Close()
	g.Body.Close()
}

func GzipHandle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.Contains(r.Header.Get("Content-Encoding"), "gzip") {
			gzipReader, err := gzip.NewReader(r.Body)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			gz := Gzip{*gzipReader, r.Body}
			defer gz.Close()

			r.Body = &gz.GzipReader
		}

		withGzNext := middleware.Compress(5)

		withGzNext(next).ServeHTTP(w, r)
	})
}
