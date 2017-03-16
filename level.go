package goseq

// Level represents the log level
type Level int

const (
	VERBOSE Level = iota
	DEBUG
	INFORMATION
	WARNING
	ERROR
	FATAL
)

var levelNames = []string{
	"VERBOSE",
	"DEBUG",
	"INFORMATION",
	"WARNING",
	"ERROR",
	"FATAL",
}

func (l Level) String() string {
	return levelNames[l]
}
