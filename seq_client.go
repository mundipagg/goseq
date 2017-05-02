package goseq

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

const (
	endpoint = "/api/events/raw"
)

// SeqClient holds the Send methods and SEQ BaseURL
type SeqClient struct {
	BaseURL string
}

// Send send POST requests to the SEQ API
func (sc *SeqClient) Send(event *SeqLog, apiKey string, client *http.Client) error {

	fullURL := sc.BaseURL + endpoint

	serialized, _ := json.Marshal(event)

	request, err := http.NewRequest("POST", fullURL, bytes.NewBuffer(serialized))

	if len(apiKey) > 1 {
		request.Header.Set("X-Seq-ApiKey", apiKey)
		request.Header.Set("Content-Type", "application/json")
	}

	if err != nil {
		log.Fatal(err)
	}
	response, err := client.Do(request)
	defer request.Body.Close()
	if err != nil {

		return err
	}
	defer response.Body.Close()

	if response.StatusCode != 201 {
		return errors.New(response.Status)
	}
	return nil
}
