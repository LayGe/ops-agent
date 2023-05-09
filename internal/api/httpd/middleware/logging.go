package middleware

import (
	"net/http"
	"time"
	"ysxs_ops_agent/pkg/log"
)

type LoggingMiddleware struct{}

func NewLoggingMiddleware() *LoggingMiddleware {
	return &LoggingMiddleware{}
}

func (m *LoggingMiddleware) Handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		start := time.Now()
		next.ServeHTTP(writer, request)
		log.MainLog.Infof(
			"[http request] proto: %s, uri: %s, method: %s, remote: %s, duration: %fs",
			request.Proto,
			request.RequestURI,
			request.Method,
			request.RemoteAddr,
			time.Since(start).Seconds())
	})
}
