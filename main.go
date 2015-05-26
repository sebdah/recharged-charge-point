package main

import (
	"flag"
	"fmt"
	"net/url"

	goLogging "github.com/op/go-logging"
	"github.com/sebdah/recharged-charge-point/config"
	"github.com/sebdah/recharged-charge-point/logging"
	"github.com/sebdah/recharged-shared/websockets"
)

var (
	log      goLogging.Logger
	WsClient *websockets.Client
)

func main() {
	// Configure logging
	logging.Setup()

	log.Info("Starting re:charged charge-point simulator")
	log.Info("Environment: %s", config.Env)

	// Parse command line options
	action := flag.String("action", "", "OCPP action")
	flag.Parse()

	// Connect to the WebSockets endpoint
	log.Info("port: %d", config.Config.GetInt("port"))
	log.Info(config.Config.GetString("central-system.endpoint-ocpp20j"))
	wsEndpoint, _ := url.Parse(config.Config.GetString("central-system.endpoint-ocpp20j"))
	log.Debug("Connecting to %s over websockets", wsEndpoint.String())
	WsClient = websockets.NewClient(wsEndpoint)

	// Start the websockets communicator
	go websocketsCommunicator()

	// Send the actions
	if *action == "authorize" {
		WsClient.WriteMessage <- `[2, "1234", "Authorize", { "idTag": { "id": "1" } }]`
	} else if *action == "bootnotification" {
		WsClient.WriteMessage <- `[2, "1234", "BootNotification", { "chargePointModel": "Model X", "chargePointVendor": "Vendor Y"}]`
	} else if *action == "datatransfer" {
		WsClient.WriteMessage <- `[2, "1234", "DataTransfer", { "vendorId": "1234" }]`
	}

	// Do not terminate
	var input string
	fmt.Scanln(&input)
	fmt.Println("done")
}

// Websockets message parser
func websocketsCommunicator() {
	var recv_msg string

	for {
		recv_msg = <-WsClient.ReadMessage

		log.Info("Received message: %s\n", recv_msg)
	}
}
