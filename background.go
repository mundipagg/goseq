package goseq

import (
	"log"
)

// Background represents a background channel that is used to send log messages to the SEQ API
type Background struct {
	ch  chan *Event
	url string
}

// NewBackground creates a new Background structure and creates a new Go Routine for the initBackground function
func NewBackground(url string) *Background {

	var a = &Background{
		ch:  make(chan *Event),
		url: url,
	}

	go a.initBackground()

	return a
}

// Background function that is responsable for sending log messages to the SEQ API
func (b *Background) initBackground() {

	var client = &SeqClient{BaseURL: b.url}

	for item := range b.ch {
		seqlog := SeqLog{
			Events: []*Event{item},
		}
		success := client.Send(&seqlog, "")

		if success != true {
			log.Fatal("shit went wrong")
		}
	}
}
