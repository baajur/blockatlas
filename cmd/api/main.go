package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas/api"
	_ "github.com/trustwallet/blockatlas/docs"
	"github.com/trustwallet/blockatlas/internal"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"github.com/trustwallet/blockatlas/platform"
)

const (
	defaultPort       = "8420"
	defaultConfigPath = "../../config.yml"
)

var (
	port, confPath string
	engine         *gin.Engine
)

func init() {
	port, confPath = internal.ParseArgs(defaultPort, defaultConfigPath)

	internal.InitConfig(confPath)
	logger.InitLogger()

	engine = internal.InitEngine(viper.GetString("gin.mode"))

	platform.Init(viper.GetStringSlice("platform"))
}

func main() {
	switch viper.GetString("rest_api") {
	case "swagger":
		api.SetupSwaggerAPI(engine)
	case "platform":
		api.SetupPlatformAPI(engine)
	default:
		api.SetupSwaggerAPI(engine)
		api.SetupPlatformAPI(engine)
	}
	internal.SetupGracefulShutdown(port, engine)
}
