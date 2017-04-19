package goseq

import (
	"errors"
	"time"
)

// Logger is the main struct that will be used to create logs
type Logger struct {
	DefinedLevel      Level
	background        *Background
	Properties        properties
	DefaultProperties properties

	baseURL string
}

// GetLogger create and returns a new Logger struct with a Background struct ready to send log messages
func GetLogger(url string, apiKey string) (*Logger, error) {

	if len(url) < 1 {
		return nil, errors.New("Invalid URL")
	}

	return &Logger{
		baseURL:           url,
		background:        NewBackground(url, apiKey),
		DefinedLevel:      0,
		Properties:        NewProperties(),
		DefaultProperties: NewProperties(),
	}, nil
}

// SetDefaultProperties sets the DefaultProperties variable
func (l *Logger) SetDefaultProperties(props map[string]interface{}) {

	for key, value := range props {

		l.DefaultProperties.AddProperty(key, value)
	}
}

// Close closes the logger background routine
func (l *Logger) Close() {
	l.background.Close()
}

func (l *Logger) log(lvl Level, message string, props properties) {

	if l.DefinedLevel != VERBOSE && l.DefinedLevel != lvl {
		return
	}

	for k, v := range l.DefaultProperties.Property {
		props.AddProperty(k, v)
	}

	entry := &Event{
		Level:           lvl.String(),
		Properties:      props.Property,
		Timestamp:       time.Now().Format("2006-01-02T15:04:05"),
		MessageTemplate: message,
	}

	l.background.ch <- entry

}

// Debug log messages with DEBUG level
func (l *Logger) Debug(message string, props properties) {
	l.log(DEBUG, message, props)
}

// Error log messages with ERROR level
func (l *Logger) Error(message string, props properties) {
	l.log(ERROR, message, props)
}

// Warning log messages with WARNING level
func (l *Logger) Warning(message string, props properties) {
	l.log(WARNING, message, props)
}

// Fatal log messages with FATAL level
func (l *Logger) Fatal(message string, props properties) {
	l.log(FATAL, message, props)
}

// Information log messages with INFORMATION level
func (l *Logger) Information(message string, props properties) {
	l.log(INFORMATION, message, props)
}
