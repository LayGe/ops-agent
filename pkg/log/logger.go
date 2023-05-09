package log

import (
	"strings"
	asyncLog "ysxs_ops_agent/pkg/log/async_log"
)

var (
	MainLog *Logger
)

type Logger struct {
	Loggerfile *asyncLog.LogFile `json:"-"`
	Level      string            `json:"loglevel"`
}

func (t *Logger) GetLevel() (level asyncLog.Priority) {
	switch strings.ToLower(t.Level) {
	case "all":
		level = asyncLog.LevelAll
	case "debug":
		level = asyncLog.LevelDebug
	case "info":
		level = asyncLog.LevelInfo
	case "warn":
		level = asyncLog.LevelWarn
	case "error":
		level = asyncLog.LevelError
	case "fatal":
		level = asyncLog.LevelFatal
	case "off":
		level = asyncLog.LevelOff
	default:
		level = asyncLog.LevelInfo
	}
	return
}

func (t *Logger) InitLog(filename string) (err error) {
	t.Loggerfile, err = asyncLog.NewLevelLog(filename, t.GetLevel())
	if err != nil {
		return
	}
	return
}

func (t *Logger) Debug(msg ...interface{}) error {
	return t.Loggerfile.Debug(msg)
}

func (t *Logger) Info(msg ...interface{}) error {
	return t.Loggerfile.Info(msg)
}

func (t *Logger) Warn(msg ...interface{}) error {
	return t.Loggerfile.Warn(msg)
}

func (t *Logger) Error(msg ...interface{}) error {
	return t.Loggerfile.Error(msg)
}

func (t *Logger) Fatal(msg ...interface{}) error {
	return t.Loggerfile.Fatal(msg)
}

func (t *Logger) Debugf(msg ...interface{}) error {
	return t.Loggerfile.Debugf(msg)
}

func (t *Logger) Infof(msg ...interface{}) error {
	return t.Loggerfile.Infof(msg)
}

func (t *Logger) Warnf(msg ...interface{}) error {
	return t.Loggerfile.Warnf(msg)
}

func (t *Logger) Errorf(msg ...interface{}) error {
	return t.Loggerfile.Errorf(msg)
}

func (t *Logger) Fatalf(msg ...interface{}) error {
	return t.Loggerfile.Fatalf(msg)
}
