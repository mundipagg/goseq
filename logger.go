package goseq

import (
	"errors"
	"net/http"
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
	APIKey            string
	Async             bool
}

// GetLogger creates and returns a new Logger struct
func GetLogger(url, apiKey string) (*Logger, error) {
	l, err := createLogger(url, apiKey, false, 0)
	return l, err
}

// GetAsyncLogger creates and returns a new Logger struct with a Background struct ready to send log messages asynchronously
func GetAsyncLogger(url, apiKey string, qtyConsumer int) (*Logger, error) {
	l, err := createLogger(url, apiKey, true, qtyConsumer)
	return l, err
}

func createLogger(url, apiKey string, async bool, qtyConsumer int) (*Logger, error) {
	if len(url) < 1 {
		return nil, errors.New("Invalid URL")
	}
	if len(apiKey) < 1 {
		return nil, errors.New("Invalid APIKey")
	}

	log := &Logger{
		baseURL:           url,
		APIKey:            apiKey,
		Async:             async,
		definedLevel:      0,
		Properties:        NewProperties(),
		DefaultProperties: NewProperties(),
	}

	if log.Async {
		err := log.newBackground(qtyConsumer)
		if err != nil {
			return nil, err
		}
	}

	return log, nil
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

func (l *Logger) log(lvl level, message string, props Properties) error {

	if l.definedLevel != Verbose && l.definedLevel != lvl {
		return errors.New("Invalid log level")
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

	if l.Async {
		l.channel <- entry
	} else {

		seqlog := seqLog{
			Events: []*event{entry},
		}

		var httpClient = &http.Client{
			Transport: &http.Transport{
				TLSHandshakeTimeout: 30 * time.Second,
			},
		}

		var logClient = &seqClient{baseURL: l.baseURL}

		err := logClient.send(&seqlog, l.APIKey, httpClient)

		if err != nil {
			return err
		}
	}

	return nil
}

// Debug log messages with DEBUG level
func (l *Logger) Debug(message string, props Properties) {
	l.log(Debug, message, props)
}

// Error log messages with ERROR level
func (l *Logger) Error(message string, props Properties) {
	l.log(Error, message, props)
}

// Warning log messages with WARNING level
func (l *Logger) Warning(message string, props Properties) {
	l.log(Warning, message, props)
}

// Fatal log messages with FATAL level
func (l *Logger) Fatal(message string, props Properties) {
	l.log(Fatal, message, props)
}

// Information log messages with INFORMATION level
func (l *Logger) Information(message string, props Properties) {
	l.log(Information, message, props)
}
