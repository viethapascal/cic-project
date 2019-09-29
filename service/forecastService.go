package service

import (
	"io/ioutil"
	"net/http"
)

func SendGetRequest(urls string) ([]byte, error) {

	resp, err := http.Get(urls)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	return body, nil
}
