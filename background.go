package goseq

import (
	"log"
)

type Background struct {
	ch  chan *Event
	url string
}

func NewBackground(url string) *Background {

	var a = &Background{
		ch:  make(chan *Event),
		url: url,
	}

	go a.initBackground()

	return a
}

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
