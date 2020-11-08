package util

import (
	"io/ioutil"
	"os"
)

const onlyMeFullAccess int64 = 0700

func AppendToFile(content string, filePath string) error {
	appendEditor := func(s string)string { return s + "\n" + content }
	return applyFileEditor(filePath, appendEditor)
}

func InsertToFileAt(content string, filePath string, lineNumber int) error {
	insertEditor := func(s string)string { return ReplaceLine(s, content, lineNumber) }
	return applyFileEditor(filePath, insertEditor)
}

func applyFileEditor(filePath string, editor func(s string)string) error {
	fileContent, err := LoadAsStr(filePath)
	if err != nil {
		return err
	}
	fileContent = editor(fileContent)
	return ioutil.WriteFile(filePath, []byte(fileContent), os.FileMode(onlyMeFullAccess))
}

func LoadAsStr(filePath string) (string, error) {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return string(content), nil
}