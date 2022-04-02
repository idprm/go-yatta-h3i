package handlers

import (
	"errors"
	"io/ioutil"
	"net/http"
)

func Postback(url string) (string, error) {

	req, err := http.NewRequest("POST", url, nil)
	req.Header.Set("Content-Type", "application/json; charset=utf8")

	if err != nil {
		return "", errors.New(err.Error())
	}

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return "", errors.New(err.Error())
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", errors.New(err.Error())
	}

	return string([]byte(body)), nil
}
