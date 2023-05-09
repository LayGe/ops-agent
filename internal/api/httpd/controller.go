package httpd

import (
	"net/http"
	"sync/atomic"
	"ysxs_ops_agent/internal/api/httpd/res"
	"ysxs_ops_agent/internal/schema"
)

func (s *Server) HealthCheck(w http.ResponseWriter, r *http.Request) {
	if atomic.LoadInt32(&healthy) == 1 {
		w.Write([]byte("ok"))
		return
	}
	w.WriteHeader(http.StatusServiceUnavailable)
}

func (s *Server) DeployNormal(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Transfer-Encoding", "chunked")

	var req = new(schema.DeployReq)
	if !bindAndCheck(w, r, req) {
		if err := s.dmService.DeployNormalWithStream(r.Context(), req, w); err != nil {
			res.ErrorResponse(w, r, err.Error())
		}
	}
	return
}
