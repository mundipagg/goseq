package goseq

type Background struct {
	ch  chan *Event
	url string
}

func NewBackground() *Background {
	return &Background{
		ch: make(chan *Event),
	}
}

func (b *Background) init_background() {

	var client = &SeqClient{BaseUrl: b.url}

	for item := range b.ch {
		seqlog := SeqLog{
			Events: []*Event{item},
		}
		client.Send(&seqlog, "")
	}
}
