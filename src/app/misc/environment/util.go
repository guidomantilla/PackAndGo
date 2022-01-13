package environment

import (
	"PackAndGo/src/app/misc/properties"
	"os"
)

const (
	CMD_PROPERTY_SOURCE_NAME = "CMD"
	OS_PROPERTY_SOURCE_NAME  = "OS"
)

func LoadEnvironment(cmdArgs *[]string) Environment {

	cmdSource := properties.NewDefaultPropertySource(CMD_PROPERTY_SOURCE_NAME, properties.NewDefaultProperties().FromArray(cmdArgs).Build())
	env := NewDefaultEnvironment().WithPropertySources(cmdSource).Build()

	osArgs := os.Environ()
	osSource := properties.NewDefaultPropertySource(OS_PROPERTY_SOURCE_NAME, properties.NewDefaultProperties().FromArray(&osArgs).Build())
	env.AppendPropertySources(osSource)

	return env
}
