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
	"Verbose",
	"Debug",
	"Information",
	"Warning",
	"Error",
	"Fatal",
}

func (l level) String() string {
	return levelNames[l]
}
