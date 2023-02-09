// Package gzip allows server to handle gzipped requests
package gzip

import (
	"compress/gzip"
	"io"
	"net/http"
	"strings"
)

// Gzip struct is designed to comply with the mighty devs' requirement to close underlying Reader
// Gzip.reader doesn't close anything by default ;(
type Gzip struct {
	GzipReader gzip.Reader
	Body       io.ReadCloser
}

// Close allows to close both gzip reader and the underlying reader
func (g *Gzip) Close() error {
	err := g.GzipReader.Close()

	if err != nil {
		return err
	}

	err = g.Body.Close()

	return err
}

// Read just allows one to read info from the reader
func (g *Gzip) Read(p []byte) (int, error) {
	return g.GzipReader.Read(p)
}

// GzipHandle returns a MW that reads the request encoding and if it is set to gzip, replaces a regular request body with a gzip-supporting implementation
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

			r.Body = &gz
		}

		next.ServeHTTP(w, r)
	})
}
