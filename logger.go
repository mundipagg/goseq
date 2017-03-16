package goseq

import (
	"time"
)

type Logger struct {
	DefinedLevel Level
	background   *Background

	base_url string
}

func GetLogger(url string) Logger {
	return Logger{
		base_url:     url,
		background:   NewBackground(url),
		DefinedLevel: 0,
	}
}

func (l *Logger) log(lvl Level, message string, props properties) {

	if l.DefinedLevel != VERBOSE && l.DefinedLevel != lvl {
		return
	}

	entry := &Event{
		Level:           lvl.String(),
		Properties:      props.Property,
		Timestamp:       time.Now().Format("2006-01-02T15:04:05"),
		MessageTemplate: message,
	}

	l.background.ch <- entry

}

func (l *Logger) Debug(message string, props properties) {
	l.log(DEBUG, message, props)
}

func (l *Logger) Error(message string, props properties) {
	l.log(ERROR, message, props)
}

func (l *Logger) Warning(message string, props properties) {
	l.log(WARNING, message, props)
}

func (l *Logger) Fatal(message string, props properties) {
	l.log(FATAL, message, props)
}

func (l *Logger) Information(message string, props properties) {
	l.log(INFORMATION, message, props)
}
