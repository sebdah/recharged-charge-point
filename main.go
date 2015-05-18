package main

import (
	"net/url"
	"os"

	"github.com/op/go-logging"
	"github.com/sebdah/recharged-charge-point/config"
	"github.com/sebdah/recharged-shared/websockets"
)

var log = logging.MustGetLogger("charge-point")

func main() {
	// Configure logging
	setupLogging()

	log.Info("Staring re:charged charge-point simulator")
	log.Info("Environment: %s", config.Env)

	// Connect to the WebSockets endpoint
	log.Info("port: %d", config.Config.GetInt("port"))
	log.Info(config.Config.GetString("central-system.endpoint-ocpp20j"))
	wsEndpoint, _ := url.Parse(config.Config.GetString("central-system.endpoint-ocpp20j"))
	log.Debug("Connecting to %s over websockets", wsEndpoint.String())
	_ = websockets.NewClient(wsEndpoint)
}

// Configure logging
func setupLogging() {
	// Create a logging backend
	backend := logging.NewLogBackend(os.Stderr, "", 0)

	// Set formatting
	format := logging.MustStringFormatter("%{color}%{time:15:04:05.000} %{shortfunc} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message}")
	backendFormatter := logging.NewBackendFormatter(backend, format)

	// Use the backends
	logging.SetBackend(backendFormatter)
}
