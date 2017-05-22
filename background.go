package goseq

import (
	"net/http"
	"sync"
	"time"
)

// background represents a background channel that is used to send log messages to the SEQ API
type background struct {
	ch     chan *event
	url    string
	apiKey string

	wg sync.WaitGroup
}

// newBackground creates a new Background structure and creates a new Go Routine for the initBackground function
func (l *Logger) newBackground(qtyConsumer int) error {
	if qtyConsumer < 1 {
		qtyConsumer = 1
	}
	var consumers []*background
	consumers = make([]*background, 0, 0)
	ch := make(chan *event)
	for i := 0; i < qtyConsumer; i++ {
		var a = &background{
			ch:     ch,
			url:    l.baseURL,
			apiKey: l.APIKey,
		}
		consumers = append(consumers, a)
		a.wg.Add(1)

		l.background = consumers
		l.channel = ch

		errCh := make(chan error)

		go func() {
			errCh <- a.initBackground()
		}()

		select {
		case err := <-errCh:
			return err
		}
	}
	return nil
}

// Background function that is responsible for sending log messages to the SEQ API
func (b *background) initBackground() error {
	var client = &seqClient{baseURL: b.url}
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
		seqlog := seqLog{
			Events: []*event{item},
		}

		err := client.send(&seqlog, b.apiKey, _client)

		if err != nil {
			return err
		}
	}
	return nil
}

// Close closes background channel and waits for the end of the go Routine
func (b *background) close() {
	b.wg.Wait()
}
