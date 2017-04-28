package goseq

import (
	"log"
	"net/http"
	"sync"
	"time"
)

// Background represents a background channel that is used to send log messages to the SEQ API
type Background struct {
	ch     chan *Event
	url    string
	apiKey string

	wg sync.WaitGroup
}

// NewBackground creates a new Background structure and creates a new Go Routine for the initBackground function
func NewBackground(url string, apiKey string, qtdConsumer int) ([]*Background, chan *Event) {
	var consumers []*Background
	consumers = make([]*Background, 0, 0)
	ch := make(chan *Event)
	for i := 0; i < qtdConsumer; i++ {
		var a = &Background{
			ch:     ch,
			url:    url,
			apiKey: apiKey,
		}
		consumers = append(consumers, a)
		a.wg.Add(1)
		go a.initBackground()
	}

	return consumers, ch
}

// Background function that is responsable for sending log messages to the SEQ API
func (b *Background) initBackground() {
	var client = &SeqClient{BaseURL: b.url}
	defer b.wg.Done()
	var _client = &http.Client{
		Transport: &http.Transport{
			TLSHandshakeTimeout: 30 * time.Second,
		},
	}
	for {
		item, ok := <-b.ch
		if !ok {
			break
		}
		seqlog := SeqLog{
			Events: []*Event{item},
		}

		success := client.Send(&seqlog, b.apiKey, _client)

		if success != true {
			log.Fatal("shit went wrong")
		}
	}
}

// Close closes background channel and waits for the end of the go Routine
func (b *Background) Close() {
	//close(b.ch)
	b.wg.Wait()
}
