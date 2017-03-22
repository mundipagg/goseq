package goseq

import (
	"log"
	"sync"
)

// Background represents a background channel that is used to send log messages to the SEQ API
type Background struct {
	ch     chan *Event
	url    string
	apiKey string

	wg sync.WaitGroup
}

// NewBackground creates a new Background structure and creates a new Go Routine for the initBackground function
func NewBackground(url string, apiKey string) *Background {

	var a = &Background{
		ch:     make(chan *Event),
		url:    url,
		apiKey: apiKey,
	}

	a.wg.Add(1)

	go a.initBackground()

	return a
}

// Background function that is responsable for sending log messages to the SEQ API
func (b *Background) initBackground() {

	var client = &SeqClient{BaseURL: b.url}

	defer b.wg.Done()

	for item := range b.ch {
		seqlog := SeqLog{
			Events: []*Event{item},
		}
		success := client.Send(&seqlog, b.apiKey)

		if success != true {
			log.Fatal("shit went wrong")
		}
	}
}

// Close closes background channel and waits for the end of the go Routine
func (b *Background) Close() {

	close(b.ch)

	b.wg.Wait()
}
