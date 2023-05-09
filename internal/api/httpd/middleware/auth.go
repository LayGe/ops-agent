package middleware

import (
	"net/http"
	"ysxs_ops_agent/pkg/log"
	"ysxs_ops_agent/utils"
	"ysxs_ops_agent/utils/ip_addr"
)

type AuthMiddleware struct {
	ipAddr string
}

func NewAuthMiddleware() *RecoverMiddleware {
	return &RecoverMiddleware{}
}

func (m *AuthMiddleware) Handler(next http.Handler) http.Handler {
	ip, err := ip_addr.GetLocalIPV4()
	if err != nil {
		log.MainLog.Errorf("failed to get ipv4 address, err:[%s]", err.Error())
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if token := r.Header.Get("token"); token == "" || token != utils.GenMD5(ip) {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("request unauthorized"))
			return
		}
		next.ServeHTTP(w, r)
	})
}
