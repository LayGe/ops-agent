package httpd

import (
	"github.com/gorilla/mux"
	"net/http"
	"ysxs_ops_agent/internal/api/httpd/middleware"
)

func (s *Server) registerWebHandlers() {
	var routers = mux.NewRouter()
	routers.HandleFunc("/health", s.HealthCheck).Methods(http.MethodGet)

	// 业务路由
	apiRoute := routers.PathPrefix("/api/v1").Subrouter()
	apiRoute.Use(
		middleware.NewLoggingMiddleware().Handler,
		middleware.NewRecoverMiddleware().Handler,
		middleware.NewAuthMiddleware().Handler,
	)

	apiRoute.HandleFunc("/deploy", s.DeployNormal).Methods(http.MethodPost)
	s.router = routers
}
