package middleware

import (
	"compress/gzip"
	"fmt"
	"net/http"
	"strings"
)

func Compression(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		compressionAlgo := r.Header.Get("Accept-Encoding")
		if !strings.Contains(compressionAlgo, "gzip") {
			next.ServeHTTP(w, r)
		}

		fmt.Printf("Compression algorithm: %s\n", compressionAlgo)
		gz := gzip.NewWriter(w)
		defer gz.Close()

		w.Header().Set("Content-Encoding", "gzip")
		w = &GzipResponseWriter{
			ResponseWriter: w,
			Writer:         gz,
		}
		next.ServeHTTP(w, r)
	})

}

type GzipResponseWriter struct {
	http.ResponseWriter
	Writer *gzip.Writer
}

func (w *GzipResponseWriter) Write(b []byte) (int, error) {
	return w.Writer.Write(b)
}