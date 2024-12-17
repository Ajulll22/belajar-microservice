package service

import (
	"bytes"
	"net/http"
	"strings"
)

func LeadToBe(requestMethod string, requestURL string, requestBytes []byte) (res *http.Response, err error) {
	method := strings.ToUpper(requestMethod)
	if method == "GET" || method == "DELETE" {

		req, err := http.NewRequest(method, requestURL, nil)
		if err != nil {
			return res, err
		}

		res, err = http.DefaultClient.Do(req)
		if err != nil {
			return res, err
		}

	} else {

		bytesToReader := bytes.NewReader(requestBytes)

		req, err := http.NewRequest(method, requestURL, bytesToReader)
		if err != nil {
			return res, err
		}

		res, err = http.DefaultClient.Do(req)
		if err != nil {
			return res, err
		}

	}

	return res, nil
}
