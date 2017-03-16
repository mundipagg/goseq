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

type properties struct {
	Property map[string]string
}

func (p *properties) AddProperty(key string, value string) {
	p.Property[key] = value
}

func NewProperties() (p properties) {
	return properties{
		Property: make(map[string]string),
	}
}
