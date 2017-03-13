package goseq

// http://localhost:5341/

import (
	"testing"
	"time"
)

func TestLogger_INFORMATION(t *testing.T) {

	logger := GetLogger("http://localhost:5341")

	logger.Information("Logging test message", nil)

	time.Sleep(100 * time.Millisecond)

}

func TestLogger_WARNING(t *testing.T) {

	logger := GetLogger("http://localhost:5341")

	logger.Warning("Logging test message", nil)

	time.Sleep(100 * time.Millisecond)

}
