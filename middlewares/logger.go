package middlewares

import (
	"log"
	"net/http"
	"time"
)

type wrappedResponseWriter struct {
	http.ResponseWriter
	StatusCode int
}

func (w *wrappedResponseWriter) WriteHeader(code int) {
	w.StatusCode = code
	w.ResponseWriter.WriteHeader(code)
}

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		wrapped := &wrappedResponseWriter{w, http.StatusOK}
		next.ServeHTTP(wrapped, r)
		log.Printf(
			"| %v | %v | %v | %v",
			wrapped.StatusCode,
			r.Method,
			r.URL.Path,
			time.Since(start),
		)
	})
}
