package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/Monska85/telegram-gateway/lib/configHelper"
	"github.com/Monska85/telegram-gateway/lib/httpHandlers"
	"github.com/Monska85/telegram-gateway/lib/logHelper"
	"github.com/Monska85/telegram-gateway/lib/telegram"
	"github.com/Monska85/telegram-gateway/lib/utils"
)

var logger = logHelper.GetInstance()

func checkRequiredEnvVars() {
	errorMessages := []string{}
	requiredEnvVars := []string{"API_TOKEN"}
	for _, envVar := range requiredEnvVars {
		if os.Getenv(envVar) == "" {
			errorMessages = append(errorMessages, fmt.Sprintf("Required environment variable '%s' is not set", envVar))
		}
	}

	if len(errorMessages) > 0 {
		for _, errorMessage := range errorMessages {
			logger.Out(utils.LogError, errorMessage)
		}
		logger.Out(utils.LogError, "Exiting due to missing environment variables")
		os.Exit(1)
	}

	logger.Out(utils.LogDebug, "All required environment variables are set")
}

func startTBotAndInitHttpHandlers() {
	var tbot *telegram.TelegramBot = telegram.GetInstance(os.Getenv("API_TOKEN"))
	httpHandlers.SetTelegramBot(tbot)
}

func main() {
	checkRequiredEnvVars()

	// Initialize the configuration helper to be sure that the configuration is properly set up.
	configHelper.GetInstance()

	// Initialize the http handlers.
	startTBotAndInitHttpHandlers()

	// Port we listen on.
	var portNum string = ":" + utils.GetEnv("PORT", "8080")
	logger.Out(utils.LogDebug, "Starting our simple http server")
	// Registering our handler functions, and creating paths.
	mux := http.NewServeMux()
	mux.HandleFunc("/sendmessage", httpHandlers.SendMessage)

	logger.Out(utils.LogDebug, "Started", "port", portNum)
	fmt.Println("To close connection CTRL+C :-)")

	// Spinning up the server.
	err := http.ListenAndServe(portNum, mux)
	if err != nil {
		logger.Out(utils.LogError, err.Error())
	}
}
