package goseq

import (
	"errors"
	"time"
)

// Logger is the main struct that will be used to create logs
type Logger struct {
	definedLevel      level
	background        []*background
	Properties        Properties
	DefaultProperties Properties
	channel           chan *event
	baseURL           string
}

// GetLogger create and returns a new Logger struct with a Background struct ready to send log messages
func GetLogger(url string, apiKey string, qtyConsumer int) (*Logger, error) {
	if len(url) < 1 {
		return nil, errors.New("Invalid URL")
	}
	lg := &Logger{
		baseURL: url,

		definedLevel:      0,
		Properties:        NewProperties(),
		DefaultProperties: NewProperties(),
	}
	backs, channel := newBackground(url, apiKey, qtyConsumer)
	lg.background = backs
	lg.channel = channel
	return lg, nil
}

// SetDefaultProperties sets the DefaultProperties variable
func (l *Logger) SetDefaultProperties(props map[string]interface{}) {

	for key, value := range props {
		l.DefaultProperties.AddProperty(key, value)
	}
}

// Close closes the logger background routine
func (l *Logger) Close() {
	close(l.channel)
	for _, back := range l.background {
		back.close()
	}

}

func (l *Logger) log(lvl level, message string, props Properties) {

	if l.definedLevel != VERBOSE && l.definedLevel != lvl {
		return
	}

	for k, v := range l.DefaultProperties.Property {
		props.AddProperty(k, v)
	}

	entry := &event{
		Level:           lvl.String(),
		Properties:      props.Property,
		Timestamp:       time.Now().Format("2006-01-02T15:04:05"),
		MessageTemplate: message,
	}
	l.channel <- entry

}

// Debug log messages with DEBUG level
func (l *Logger) Debug(message string, props Properties) {
	l.log(DEBUG, message, props)
}

// Error log messages with ERROR level
func (l *Logger) Error(message string, props Properties) {
	l.log(ERROR, message, props)
}

// Warning log messages with WARNING level
func (l *Logger) Warning(message string, props Properties) {
	l.log(WARNING, message, props)
}

// Fatal log messages with FATAL level
func (l *Logger) Fatal(message string, props Properties) {
	l.log(FATAL, message, props)
}

// Information log messages with INFORMATION level
func (l *Logger) Information(message string, props Properties) {
	l.log(INFORMATION, message, props)
}
