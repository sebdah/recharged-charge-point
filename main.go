package main

import (
	"net/url"

	goLogging "github.com/op/go-logging"
	"github.com/sebdah/recharged-charge-point/config"
	"github.com/sebdah/recharged-charge-point/logging"
	"github.com/sebdah/recharged-shared/websockets"
)

var log goLogging.Logger

func main() {
	// Configure logging
	logging.Setup()

	log.Info("Staring re:charged charge-point simulator")
	log.Info("Environment: %s", config.Env)

	// Connect to the WebSockets endpoint
	log.Info("port: %d", config.Config.GetInt("port"))
	log.Info(config.Config.GetString("central-system.endpoint-ocpp20j"))
	wsEndpoint, _ := url.Parse(config.Config.GetString("central-system.endpoint-ocpp20j"))
	log.Debug("Connecting to %s over websockets", wsEndpoint.String())
	_ = websockets.NewClient(wsEndpoint)
}
