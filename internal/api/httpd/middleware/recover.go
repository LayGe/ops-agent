package middleware

import (
	"net/http"
	"ysxs_ops_agent/pkg/log"
)

type RecoverMiddleware struct{}

func NewRecoverMiddleware() *RecoverMiddleware {
	return &RecoverMiddleware{}
}

func (m *RecoverMiddleware) Handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		defer func() {
			if p := recover(); p != nil {
				log.MainLog.Errorf(
					"proto: %s, remote_addr: %s, method: %s, uri: %s, request raise error: ",
					request.Proto,
					request.RemoteAddr,
					request.Method,
					request.RequestURI,
					p)
				writer.WriteHeader(http.StatusInternalServerError)
				return
			}
		}()
		next.ServeHTTP(writer, request)
	})
}
