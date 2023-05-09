package app

import (
	"context"
	"log"
	"path/filepath"
	"time"
	"ysxs_ops_agent/config"
	"ysxs_ops_agent/internal/api/httpd"
	"ysxs_ops_agent/internal/service"
	agent_log "ysxs_ops_agent/pkg/log"
	"ysxs_ops_agent/pkg/signals"
)

var Version = "unknown"

func RunAgentForever(confPath string) {
	config.Setup(confPath)

	// init logger
	agent_log.MainLog = new(agent_log.Logger)
	agent_log.MainLog.Level = config.GetConf().LogLevel
	logPath := filepath.Join(config.GetConf().LogDirPath, "ops_agent.log")
	err := agent_log.MainLog.InitLog(logPath)
	if err != nil {
		log.Fatalf("init MainLog failed, err:[%s]", err.Error())
	}
	// service层
	dmsService := service.NewDMService()

	// controller层
	// 初始化httpd
	httpdSrv := httpd.NewServer(dmsService)
	ctx := context.Background()
	httpServer, healthy := httpdSrv.ListenAndServe(ctx)

	// graceful shutdown
	stopCh := signals.SetupSignalHandler()
	sd, _ := signals.NewShutdown(30 * time.Second)
	sd.Graceful(stopCh, httpServer, healthy)
}
