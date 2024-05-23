package configHelper

import (
	"os"
	"sync"

	"github.com/Monska85/telegram-gateway/lib/logHelper"
	"github.com/Monska85/telegram-gateway/lib/utils"
	"gopkg.in/yaml.v3"
)

var logger = logHelper.GetInstance()

type configuration struct {
	MessageRouter struct {
		Default struct {
			Patterns []map[string]string `yaml:"patterns,omitempty"`
		} `yaml:"default,omitempty"`
	} `yaml:"messageRouter,omitempty"`
}

type ConfigHelper struct {
	Config *configuration
}

var instance *ConfigHelper
var lock = &sync.Mutex{}

func GetInstance() *ConfigHelper {
	if instance != nil {
		logger.Out(utils.LogDebug, "Single instance of ConfigHelper already created")
		return instance
	}

	lock.Lock()
	defer lock.Unlock()

	if instance != nil {
		logger.Out(utils.LogDebug, "Single instance of ConfigHelper already created")
		return instance
	}

	logger.Out(utils.LogDebug, "Creating new ConfigHelper instance")

	var config *configuration
	var configurationFilePath = utils.GetEnv("CONFIGURATION_FILE_PATH", utils.ConfigurationFilePath)

	// Check if the configuration file exists. If it does not, we will use the default configuration.
	_, err := os.Stat(configurationFilePath)
	if os.IsNotExist(err) {
		instance = &ConfigHelper{Config: config}
		return instance
	}

	buf, err := os.ReadFile(configurationFilePath)
	if err != nil {
		logger.Out(utils.LogError, "Error reading configuration file", "error", err.Error())
		os.Exit(1)
	}

	config = &configuration{}
	err = yaml.Unmarshal(buf, config)
	if err != nil {
		logger.Out(utils.LogError, "Error unmarshalling configuration file", "error", err.Error())
		os.Exit(1)
	}

	logger.Out(utils.LogDebug, "Configuration file loaded successfully", "config", config)

	instance = &ConfigHelper{Config: config}
	return instance
}
