package service

import (
	"api-listeners/app/util"
	"log"
	"strconv"
	"strings"
)

type ConfigService interface {
	GetProp(key string) string
	GetIntProp(key string) int64
}

type EmbeddedFileConfigService struct {
	FilePath string
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
		service.loadConfig()
	}
	return service.props
}

func (service *EmbeddedFileConfigService) loadConfig() {
	configStr, err := util.LoadAsStr(service.FilePath)
	if err != nil {
		log.Fatal(err)
	}
	props := make(map[string]string)
	for _, line := range strings.Split(configStr, "\n") {
		if isValidConfigLine(line) {
			keyAndValue := strings.Split(line, "=")
			props[keyAndValue[0]] = keyAndValue[1]
		}
	}
	service.props = props
}

func isValidConfigLine(l string) bool {
	return !shouldIgnoreConfigLine(l)
}

func shouldIgnoreConfigLine(l string) bool {
	lNoSpaces := strings.TrimSpace(l)
	return strings.HasPrefix(lNoSpaces, "#") || len(lNoSpaces) == 0
}