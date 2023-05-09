package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sync"
)

const (
	hostEnvKey = "SERVER_HOSTNAME"
)

var (
	GlobalConfig *Config
	lock         = new(sync.RWMutex)
)

type Config struct {
	AgentName   string `yaml:"agent_name"`
	Debug       bool   `yaml:"debug"`
	BindAddress string `yaml:"bind_address"`
	LogDirPath  string
	LogLevel    string `yaml:"log_level"`
}

func GetConf() Config {
	if GlobalConfig == nil {
		return getDefaultConfig()
	}
	return *GlobalConfig
}

func Setup(confPath string) {
	lock.RLock()
	defer lock.RUnlock()

	defaultConfig := getDefaultConfig()
	loadConfigFromFile(confPath, &defaultConfig)
	GlobalConfig = &defaultConfig
	log.Printf("%+v\n", GlobalConfig)
}

func getDefaultConfig() Config {
	defaultName := getDefaultName()
	rootPath := getPwdDirPath()
	logDirPath := filepath.Join(rootPath, "logs")

	folders := []string{logDirPath}
	for _, folder := range folders {
		if err := ensureDirExist(folder); err != nil {
			log.Fatalf("create file failed, err:[%s]", err.Error())
		}
	}

	return Config{
		AgentName:   defaultName,
		Debug:       true,
		BindAddress: "0.0.0.0:18088",
		LogDirPath:  logDirPath,
		LogLevel:    "INFO",
	}
}

func loadConfigFromFile(confPath string, conf *Config) {
	if have(confPath) {
		cfgFile, err := ioutil.ReadFile(confPath)
		if err == nil {
			if err = yaml.Unmarshal(cfgFile, conf); err == nil {
				log.Printf("load config from %s success\n", confPath)
			}
		}
		if err != nil {
			log.Fatalf("load config from %s failed: %s\n", confPath, err.Error())
		}
	}
}

/*
SERVER_HOSTNAME: 环境变量名，用于自定义默认注册名称的前缀
default name rule:
{SERVER_HOSTNAME}-{HOSTNAME}
 or
{HOSTNAME}
*/
func getDefaultName() string {
	hostname, _ := os.Hostname()
	if serverHostname, ok := os.LookupEnv(hostEnvKey); ok {
		hostname = fmt.Sprintf("%s-%s", serverHostname, hostname)
	}
	return hostname
}

func getPwdDirPath() string {
	if rootPath, err := os.Getwd(); err == nil {
		return rootPath
	}
	return ""
}

func have(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}
func haveDir(file string) bool {
	fi, err := os.Stat(file)
	return err == nil && fi.IsDir()
}

func ensureDirExist(file string) error {
	if !haveDir(file) {
		if err := os.MkdirAll(file, os.ModePerm); err != nil {
			return err
		}
	}
	return nil
}
