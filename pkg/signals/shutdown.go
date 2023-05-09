package signals

import (
	"context"
	"net/http"
	"sync/atomic"
	"time"
	"ysxs_ops_agent/config"
	"ysxs_ops_agent/pkg/log"
)

type Shutdown struct {
	serverShutdownTimeout time.Duration
}

func NewShutdown(serverShutdownTimeout time.Duration) (*Shutdown, error) {
	srv := &Shutdown{
		serverShutdownTimeout: serverShutdownTimeout,
	}
	return srv, nil
}

func (s *Shutdown) Graceful(stopCh <-chan struct{}, httpServer *http.Server, healthy *int32) {
	ctx := context.Background()

	<-stopCh
	ctx, cancel := context.WithTimeout(ctx, s.serverShutdownTimeout)
	defer cancel()

	atomic.StoreInt32(healthy, 0)

	log.MainLog.Info("Shutting down HTTP/HTTPS server")

	if config.GetConf().Debug {
		time.Sleep(3 * time.Second)
	}

	// determine if the secure server was started
	if httpServer != nil {
		if err := httpServer.Shutdown(ctx); err != nil {
			log.MainLog.Warn("HTTPS server graceful shutdown failed")
		}
	}
}
