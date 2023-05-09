package httpd

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"sync/atomic"
	"time"
	"ysxs_ops_agent/config"
	"ysxs_ops_agent/internal/service"
	"ysxs_ops_agent/pkg/log"
)

var (
	healthy int32
)

type Server struct {
	router    *mux.Router
	Srv       *http.Server
	handler   http.Handler
	dmService *service.DMService
}

func NewServer(dmService *service.DMService) *Server {
	srv := &Server{
		router:    mux.NewRouter(),
		dmService: dmService,
	}
	return srv
}

func (s *Server) ListenAndServe(ctx context.Context) (*http.Server, *int32) {
	s.registerWebHandlers()

	s.handler = s.router
	srv := s.startServer(ctx)
	atomic.StoreInt32(&healthy, 1)

	return srv, &healthy
}

func (s *Server) startServer(_ context.Context) *http.Server {
	fmt.Println(s.handler)
	srv := &http.Server{
		Addr:         config.GetConf().BindAddress,
		WriteTimeout: time.Minute * 10,
		ReadTimeout:  time.Minute * 10,
		IdleTimeout:  time.Second * 10,
		Handler:      s.handler,
	}
	//s.printRoutes()

	// 启动 http server
	go func() {
		log.MainLog.Infof("Starting http server on %s.", srv.Addr)
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			log.MainLog.Fatal("HTTP server crashed")
		}
	}()
	return srv
}
