package goseq

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

const (
	ENDPOINT = "/api/events/raw"
)

type SeqClient struct {
	BaseUrl string
}

func (sc *SeqClient) Send(event Event, api_key string) {

	fullUrl := sc.BaseUrl + ENDPOINT

	serialized, _ := json.Marshal(event)

	request, err := http.NewRequest("POST", fullUrl, bytes.NewBuffer(serialized))

	if len(api_key) > 1 {
		request.Header.Set("X-Seq-ApiKey", api_key)
	}

	if err != nil {
		log.Fatal(err)
	}

	client := &http.Client{}

	response, err := client.Do(request)

	defer response.Body.Close()
}
