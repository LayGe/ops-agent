package main

import (
	"flag"
	"fmt"
	"ysxs_ops_agent/cmd/agent/app"
)

var (
	//pidPath = "/tmp/ops_agent.pid"
	//version       = "unknown"
	//daemonFlag    = false
	runSignalFlag = "start"
	infoFlag      = false
	configPath    = ""
)

func init() {
	//flag.BoolVar(&daemonFlag, "d", false, "start as daemon(bool)")
	flag.StringVar(&runSignalFlag, "s", "start", "start | stop")
	flag.StringVar(&configPath, "c", "ops_agent.yml", "ops_agent.yml path")
	flag.BoolVar(&infoFlag, "v", false, "version info(bool)")
}

//func startAsDaemon() {
//	ctx := &daemon.Context{
//		PidFileName: "ops_agent.pid",
//		PidFilePerm: 0644,
//		Umask:       027,
//		WorkDir:     "./",
//	}
//	child, err := ctx.Reborn()
//	if err != nil {
//		log.Fatalf("run failed: %v", err)
//	}
//	if child != nil {
//		return
//	}
//	defer ctx.Release()
//	app.RunAgentForever(configPath)
//}

func main() {
	flag.Parse()
	if infoFlag {
		fmt.Printf("Version:   %s\n", app.Version)
		return
	}
	//if runSignalFlag == "stop" {
	//	pid, err := ioutil.ReadFile(pidPath)
	//	if err != nil {
	//		log.Fatalf("pid file: %s not exist", pidPath)
	//		return
	//	}
	//	pidInt, _ := strconv.Atoi(string(pid))
	//	err = syscall.Kill(pidInt, syscall.SIGTERM)
	//	if err != nil {
	//		log.Fatalf("Stop failed: %v", err)
	//	} else {
	//		_ = os.Remove(pidPath)
	//	}
	//	return
	//}
	app.RunAgentForever(configPath)
	//switch {
	//case daemonFlag:
	//	startAsDaemon()
	//default:
	//	app.RunAgentForever(configPath)
	//}
}
