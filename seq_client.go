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

// seqClient holds the Send methods and SEQ BaseURL
type seqClient struct {
	baseURL string
}

// send send POST requests to the SEQ API
func (sc *seqClient) send(event *seqLog, apiKey string, client *http.Client) error {
	fullURL := sc.baseURL + endpoint

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
