package goseq

// Event represents an log entry on SEQ
type Event struct {
	Timestamp       string
	Level           string
	MessageTemplate string
	Properties      map[string]interface{}
}

// SeqLog is Event array
type SeqLog struct {
	Events []*Event
}

type properties struct {
	Property map[string]interface{}
}

func (p *properties) AddProperty(key string, value interface{}) {
	p.Property[key] = value
}

// NewProperties creates a new properties struct and creates a new Property Map
func NewProperties() (p properties) {
	return properties{
		Property: make(map[string]interface{}),
	}
}
