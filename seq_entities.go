package goseq

// Event represents an log entry on SEQ
type Event struct {
	Timestamp       string
	Level           string
	MessageTemplate string
	Properties      map[string]string
}

// SeqLog is Event array
type SeqLog struct {
	Events []*Event
}

type properties struct {
	Property map[string]string
}

func (p *properties) AddProperty(key string, value string) {
	p.Property[key] = value
}

// NewProperties creates a new properties struct and creates a new Property Map
func NewProperties() (p properties) {
	return properties{
		Property: make(map[string]string),
	}
}
