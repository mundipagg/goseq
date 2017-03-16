## Golang logging library for SEQ tool

### [SEQ](https://getseq.net/)

Structured logs for .NET apps

Seq is the fastest way for development teams to carry the benefits of structured logging from development through to production.

Modern structured logging bridges the gap between human-friendly text logs, and machine-readable formats like JSON. Using event data from libraries such as Serilog, ASP.NET Core, and Node.js, Seq makes centralized logs easy to read, and easy to filter and correlate, without fragile log parsing.


### Exemples

```go
package goseq

// http://localhost:5341/

import (
	"testing"
	"time"
)

func TestLogger_INFORMATION(t *testing.T) {

	logger := GetLogger("http://localhost:5341")

	logger.Information("Logging test message", NewProperties())

	time.Sleep(100 * time.Millisecond)

}
```