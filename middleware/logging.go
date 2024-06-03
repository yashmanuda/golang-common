package middleware

import (
	"net/http"
	"time"

	"github.com/yashmanuda/golang-common/logger"
)

var log = logger.GetLoggerInstance()

type WrappedResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (rw *WrappedResponseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		wrappedWriter := &WrappedResponseWriter{ResponseWriter: w}
		log.Info(
			"Request started",
			"method", r.Method,
			"url", r.URL.Path,
		)
		next.ServeHTTP(wrappedWriter, r)
		log.Info(
			"Request completed",
			"method", r.Method,
			"url", r.URL.Path,
			"duration", time.Since(start).Milliseconds(),
			"status", wrappedWriter.statusCode,
		)
	})
}
