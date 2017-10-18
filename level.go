package goseq

// Level represents the log level
type level int

//Log level supported by Seq
const (
	Verbose level = iota
	Debug
	Information
	Warning
	Error
	Fatal
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
