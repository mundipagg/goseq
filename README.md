# Golang logging library for SEQ tool

## GoSeq

GoSeq package implements a logging infrascruture that enables to use SEQ logging tool. 


### [SEQ](https://getseq.net/)

Structured logs for .NET apps

Seq is the fastest way for development teams to carry the benefits of structured logging from development through to production.

Modern structured logging bridges the gap between human-friendly text logs, and machine-readable formats like JSON. Using event data from libraries such as Serilog, ASP.NET Core, and Node.js, Seq makes centralized logs easy to read, and easy to filter and correlate, without fragile log parsing.


## Examples

```go
package goseq

// SEQ URL http://localhost:5341/

import (
	"testing"
	"time"
)

func TestLogger_INFORMATION(t *testing.T) {

	// Creates a new logger instance that will be used to send log messages to SEQ API
	logger := GetLogger("http://localhost:5341")

	// Logs message with INFORMATION log level and empty properties
	logger.Information("Logging test message", NewProperties())
	
	// Closes logger inner channel and end the Go Routine responsably for logging messages.
	// Close MUST always be called at the end of the application to avoid loosing log messages
	logger.Close()

}
```

## Installing

### Use *go get*

    $ go get github.com/mundipagg/goseq

Its source will be in:

    $GOPATH/src/pkg/github.com/mundipagg/goseq

You can use `go get -u` to update the package.