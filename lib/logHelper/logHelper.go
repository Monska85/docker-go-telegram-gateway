package logHelper

import (
	"log/slog"
	"os"
	"reflect"
	"strings"
	"sync"

	"github.com/Monska85/telegram-gateway/lib/utils"
)

type LogHelper struct {
	logger *slog.Logger
}

var instance *LogHelper
var lock = &sync.Mutex{}

func (l LogHelper) Out(loggerMethod string, message string, args ...any) {
	r := reflect.TypeOf(l.logger)
	_, ok := r.MethodByName(loggerMethod)
	if !ok {
		l.logger.Error("Method not found", "method", loggerMethod, "args", args)
		return
	}

	var finalArgs = []reflect.Value{reflect.ValueOf(message)}
	if len(args) > 0 {
		for _, arg := range args {
			finalArgs = append(finalArgs, reflect.ValueOf(arg))
		}
	}
	reflect.ValueOf(l.logger).MethodByName(loggerMethod).Call(finalArgs)
}

func GetInstance() *LogHelper {
	if instance != nil {
		instance.Out(utils.LogDebug, "Single instance of LogHelper already created")
		return instance
	}

	lock.Lock()
	defer lock.Unlock()

	if instance != nil {
		instance.Out(utils.LogDebug, "Single instance of LogHelper already created")
		return instance
	}

	var logLevel = utils.GetEnv("LOG_LEVEL", "info")
	var slogOptions = &slog.HandlerOptions{}
	if strings.EqualFold(logLevel, "debug") {
		slogOptions.Level = slog.LevelDebug
	}

	instance = &LogHelper{logger: slog.New(slog.NewJSONHandler(os.Stdout, slogOptions))}
	instance.Out(utils.LogDebug, "Creating new LogHelper instance", "logLevel", logLevel)
	return instance
}
