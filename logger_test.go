package goseq

// http://localhost:5341/

import (
	"fmt"
	"testing"
)

func TestLogger_INFORMATION(t *testing.T) {

	logger, _ := GetLogger("http://localhost:5341")

	logger.Information("Logging test message", NewProperties())

}

func TestLogger_WARNING(t *testing.T) {

	logger, _ := GetLogger("http://localhost:5341")

	logger.Warning("Logging test message", NewProperties())

}

func TestLogger_WithArgs(t *testing.T) {

	logger, _ := GetLogger("http://localhost:5341")

	var props = NewProperties()
	props.AddProperty("GUID", "11AE3484-9CD4-4332-98B1-145AAEBEACAB")
	props.AddProperty("String", "SEQ")
	props.AddProperty("Key", "Value")

	logger.Warning("Message with args", props)

}

func BenchmarkLogger_WithArgs_100times(b *testing.B) {

	logger, _ := GetLogger("http://localhost:5341")

	var props = NewProperties()
	props.AddProperty("GUID", "11AE3484-9CD4-4332-98B1-145AAEBEACAB")
	props.AddProperty("String", "SEQ")
	props.AddProperty("Key", "Value")

	for index := 0; index < 100; index++ {
		logger.Warning(fmt.Sprintf("Message with args %d", index), props)
	}

	logger.Close()

}

func TestLogger_URLError(t *testing.T) {

	_, err := GetLogger("")

	if err != nil {
		t.Log("Worked")
	}

}

func TestLogger_URLError_Fail(t *testing.T) {

	_, err := GetLogger("")

	if err == nil {
		t.FailNow()
	}
}
