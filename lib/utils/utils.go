package utils

import (
	"encoding/json"
	"os"
	"reflect"
	"runtime"
)

func GetEnv(key string, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

func GetJsonString(obj interface{}) string {
	s, _ := json.Marshal(obj)
	return string(s)
}

func GetFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}
