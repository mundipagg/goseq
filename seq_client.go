package goseq

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

const (
	endpoint = "/api/events/raw"
)

type SeqClient struct {
	BaseURL string
}

func (sc *SeqClient) Send(event *SeqLog, api_key string) bool {

	fullURL := sc.BaseURL + endpoint

	serialized, _ := json.Marshal(event)

	request, err := http.NewRequest("POST", fullURL, bytes.NewBuffer(serialized))

	if len(api_key) > 1 {
		request.Header.Set("X-Seq-ApiKey", api_key)
		request.Header.Set("Content-Type", "application/json")
	}

	if err != nil {
		log.Fatal(err)
	}

	client := &http.Client{}

	response, err := client.Do(request)

	defer response.Body.Close()

	if response.StatusCode == 201 {
		return true
	}

	return false
}
