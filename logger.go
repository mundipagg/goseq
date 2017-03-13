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
		background:   NewBackground(),
		DefinedLevel: 0,
	}
}

func (l *Logger) log(lvl Level, message string, args map[string]string) {

	if l.DefinedLevel != lvl {
		return
	}

	entry := &Event{
		Level:      lvl.String(),
		Properties: args,
		Timestamp:  time.Now().String(),
	}

	l.background.ch <- entry

}

func (l *Logger) Debug(message string, args map[string]string) {
	l.log(DEBUG, message, args)
}

func (l *Logger) Error(message string, args map[string]string) {
	l.log(ERROR, message, args)
}

func (l *Logger) Warning(message string, args map[string]string) {
	l.log(WARNING, message, args)
}

func (l *Logger) Fatal(message string, args map[string]string) {
	l.log(FATAL, message, args)
}

func (l *Logger) Information(message string, args map[string]string) {
	l.log(INFORMATION, message, args)
}
