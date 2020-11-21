package util

import (
	"bytes"
	"encoding/json"
	"io"
)

func ToJson(content interface{}) (io.Reader, error) {
	jsonBytes, err := json.Marshal(content)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(jsonBytes), nil
}

func ParseJsonBytes(content []byte, resultJson interface{}) error {
	content = removeZeros(content)
	contentReader := bytes.NewReader(content)
	return parseJson(contentReader, resultJson)
}

func parseJson(r io.Reader, resultJson interface{}) error {
	d := json.NewDecoder(r)
	return d.Decode(resultJson)
}
