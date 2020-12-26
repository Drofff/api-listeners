package util

import (
	"fmt"
	"io"
	"net/http"
)

func DoGetWithToken(url string, token string, respBody interface{}) error {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Token", token)
	content, err := execClientRequest(req)
	if err != nil {
		return err
	}
	return ParseJsonBytes(content, respBody)
}

func execClientRequest(req *http.Request) ([]byte, error) {
	c := http.Client{}
	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	return readAll(resp)
}

func DoPostJson(url string, body interface{}, respBody interface{}) error {
	jsonBody, err := ToJson(body)
	if err != nil {
		return err
	}
	respBytes, err := DoPost(url, "application/json", jsonBody)
	if err != nil {
		return err
	}
	return ParseJsonBytes(respBytes, respBody)
}

func DoPost(url string, contentType string, body io.Reader) ([]byte, error) {
	resp, err := http.Post(url, contentType, body)
	if err != nil {
		return nil, err
	}
	if isNotSuccessCode(resp.StatusCode) {
		return nil, HttpRequestError(resp.StatusCode)
	}
	return readAll(resp)
}

func isNotSuccessCode(code int) bool {
	return code < 200 || code >= 300
}

func readAll(resp *http.Response) ([]byte, error) {
	respBody := make([]byte, 100)
	for {
		bodyPart := make([]byte, 100)
		_, err := resp.Body.Read(bodyPart)
		if err != nil {
			if err != io.EOF {
				return nil, err
			}
			return respBody, nil
		}
		respBody = append(respBody, bodyPart...)
	}
}

type HttpRequestError int

func (err HttpRequestError) Error() string {
	return fmt.Sprintf("HTTP server respond with status %v\n", int(err))
}