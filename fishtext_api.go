package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

const (
	BaseAPI = "https://fish-text.ru/get?"
)

type FishTextResponse struct {
	Status string `json:"status"`
	Text string `json:"text"`
	ErrorCode int `json:"errorCode"`
}

func getContent(query string, count int) (string, error) {
	switch query {
	case "sentence", "title":
		if count < 1 || count > 500 {
			return "", errors.New("sentences count must be 1..500")
		}
	case "paragraph":
		if count < 1 || count > 100 {
			return "", errors.New("paragraphs count must be 1..100")
		}
	}

	url := BaseAPI + "&type=" + query + "&number=" + strconv.Itoa(count)
	response, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	var fishTextResponse FishTextResponse
	err = json.Unmarshal(body, &fishTextResponse)
	if err != nil {
		return "", err
	}

	if fishTextResponse.Status != "success" {
		return "", errors.New("fishtext API error: " + strconv.Itoa(fishTextResponse.ErrorCode))
	}

	return strings.ReplaceAll(fishTextResponse.Text, "\\n\\n", "\n"), nil
}