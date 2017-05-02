package goseq

// Level represents the log level
type level int

//Log level supported by Seq
const (
	VERBOSE level = iota
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

func (l level) String() string {
	return levelNames[l]
}
