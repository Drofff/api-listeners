package service

import (
	"api-listeners/app/util"
	"fmt"
	"log"
	"strconv"
	"strings"
)

type StateHolder interface {
	GetIntState(key string) (int64, error)
	SetIntState(key string, state int64)
}

type FileStateHolder struct {
	FilePath string
}

const keyAndStateDelimiter = "="

func (service FileStateHolder) GetIntState(key string) (int64, error) {
	prefix := key + keyAndStateDelimiter
	stateStr, err := service.getStateFromFileByPrefix(prefix)
	if err != nil {
		return -1, err
	}
	state, err := strconv.ParseInt(stateStr, 0, 64)
	if err != nil {
		return -1, err
	}
	return state, nil
}

func (service FileStateHolder) getStateFromFileByPrefix(prefix string) (string, error) {
	ln, l := service.getLineWithPrefix(prefix)
	if ln == -1 {
		return "", UnknownPrefixError(prefix)
	}
	keyAndState := strings.Split(l, keyAndStateDelimiter)
	return keyAndState[1], nil
}

func (service FileStateHolder) SetIntState(key string, state int64) {
	prefix := key + keyAndStateDelimiter
	l := prefix + fmt.Sprint(state)
	ln, _ := service.getLineWithPrefix(prefix)
	var err error
	if ln == -1 {
		err = util.AppendToFile(l, service.FilePath)
	} else {
		err = util.InsertToFileAt(l, service.FilePath, ln)
	}
	if err != nil {
		log.Fatal(err)
	}
}

func (service FileStateHolder) getLineWithPrefix(prefix string) (int, string) {
	fileContent, err := util.LoadAsStr(service.FilePath)
	if err != nil {
		log.Fatal(err)
	}
	for ln, l := range strings.Split(fileContent, "\n") {
		if strings.HasPrefix(l, prefix) {
			return ln, l
		}
	}
	return -1, ""
}

type UnknownPrefixError string

func (err UnknownPrefixError) Error() string {
	return "State file has no data for prefix '" + string(err) + "'"
}