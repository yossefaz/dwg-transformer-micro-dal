package log

import (
	"fmt"

	"github.com/yossefaz/dwg-transformer-micro-utils/logs"
	"github.com/yossefaz/dwg-transformer-micro-utils/utils"
	"github.com/tkanos/gonfig"
)

var Logger utils.Logger

var configEnv = map[string]string{
	"dev":  "./config/logs.json",
	"prod": "./config/logs.json",
}

type logConfiguration struct {
	Logs struct {
		Main struct {
			Path  string
			Level string
		}
	}
}

func GetLogger(env string) {
	configuration := logConfiguration{}
	err := gonfig.GetConf(configEnv[env], &configuration)
	if err != nil {
		fmt.Println("Cannot read config file")
	}
	Logger, err = logs.InitLogs(configuration.Logs.Main.Path, configuration.Logs.Main.Level)
	if err != nil {
		fmt.Println("Cannot instantiate logger : ", err)
	}
}