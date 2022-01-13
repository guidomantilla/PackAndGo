package config

import (
	"PackAndGo/src/app/misc/environment"
	"fmt"

	"go.uber.org/zap"
)

var singletonEnvironment environment.Environment

func StopParams() {
	//Nothing to do here yet
}

func InitParams(cmdArgs *[]string) environment.Environment {
	singletonEnvironment = environment.LoadEnvironment(cmdArgs)
	for _, source := range singletonEnvironment.GetPropertySources() {
		sourceMap := source.AsMap()
		name, internalMap := sourceMap["name"], sourceMap["value"].(map[string]string)
		for key, value := range internalMap {
			zap.L().Debug(fmt.Sprintf("source name: %s, key: %s, value: %s", name, key, value))
		}
	}
	return singletonEnvironment
}
