package goseq

// event represents an log entry on SEQ
type event struct {
	Timestamp       string
	Level           string
	MessageTemplate string
	Properties      map[string]interface{}
}

// seqLog is Event array
type seqLog struct {
	Events []*event
}

// Properties used to set properties of the log
type Properties struct {
	Property map[string]interface{}
}

//AddProperty adds new property to log
func (p *Properties) AddProperty(key string, value interface{}) {
	p.Property[key] = value
}

// NewProperties creates a new properties struct and creates a new Property Map
func NewProperties() (p Properties) {
	return Properties{
		Property: make(map[string]interface{}),
	}
}
