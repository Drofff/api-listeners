package service

import (
	"api-listeners/app/util"
	"log"
	"os"
	"strconv"
	"strings"
)

type ConfigService interface {
	GetProp(key string) string
	GetIntProp(key string) int64
}

type EmbeddedFileConfigService struct {
	ConfigFilePath string
	props map[string]string
}

func (service *EmbeddedFileConfigService) GetIntProp(key string) int64 {
	propStr := service.GetProp(key)
	prop, err := strconv.ParseInt(propStr, 0, 64)
	if err != nil {
		panic(err)
	}
	return prop
}

func (service *EmbeddedFileConfigService) GetProp(key string) string {
	return service.getConfig()[key]
}

func (service *EmbeddedFileConfigService) getConfig() map[string]string {
	if service.props == nil {
		service.props = make(map[string]string)
		service.loadConfigFromFile()
		service.loadEnvConfig()
	}
	return service.props
}

func (service *EmbeddedFileConfigService) loadConfigFromFile() {
	configFileText, err := util.LoadAsStr(service.ConfigFilePath)
	if err != nil {
		log.Fatal(err)
	}
	for _, line := range strings.Split(configFileText, "\n") {
		if isValidConfigLine(line) {
			service.parseProp(line)
		}
	}
}

func isValidConfigLine(l string) bool {
	return !shouldIgnoreConfigLine(l)
}

func shouldIgnoreConfigLine(l string) bool {
	lNoSpaces := strings.TrimSpace(l)
	return strings.HasPrefix(lNoSpaces, "#") || len(lNoSpaces) == 0
}

const envKey = "ozzy_rate"

/*
	Loads configuration properties from environment variable specified by
	envKey and expects value to be of format: key_0=value_0;key_N=value_N
 */
func (service *EmbeddedFileConfigService) loadEnvConfig() {
	envProps := os.Getenv(envKey)
	if util.IsBlank(envProps) {
		return
	}
	for _, envProp := range strings.Split(envProps, ";") {
		service.parseProp(envProp)
	}
}

func (service *EmbeddedFileConfigService) parseProp(prop string) {
	keyAndValue := strings.Split(prop, "=")
	service.props[keyAndValue[0]] = keyAndValue[1]
}
