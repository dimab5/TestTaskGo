package middleware

import (
	"log"
	"net/http"
	"time"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		log.Printf("Started %s %s", r.Method, r.URL.Path)

		next.ServeHTTP(w, r)

		statusCode := http.StatusOK
		if r := w.(interface {
			Status() int
		}); r != nil {
			statusCode = r.Status()
		}

		log.Printf(
			"Completed %s %s %v in %v",
			r.Method,
			r.URL.Path,
			statusCode,
			time.Since(start),
		)
	})
}
