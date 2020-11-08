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
	content, err := removeLeadingAndTrailingZeros(content)
	if err != nil {
		return err
	}
	contentReader := bytes.NewReader(content)
	return ParseJson(contentReader, resultJson)
}

func ParseJson(r io.Reader, resultJson interface{}) error {
	d := json.NewDecoder(r)
	return d.Decode(resultJson)
}