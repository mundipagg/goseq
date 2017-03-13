package goseq

type Event struct {
	Timestamp       string
	Level           string
	MessageTemplate string
	Properties      map[string]string
}

type SeqLog struct {
	Events []*Event
}
