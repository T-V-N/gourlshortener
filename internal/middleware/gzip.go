package gzip

import (
	"net/http"
	"strings"

	"github.com/NYTimes/gziphandler"
)

func gzipHandle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			next.ServeHTTP(w, r)
			return
		}

		withGzNext := gziphandler.GzipHandler(next)

		withGzNext.ServeHTTP(w, r)
	})
}
