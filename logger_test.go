package goseq

// http://localhost:5341/

import (
	"fmt"
	"testing"
)

var baseURL = "http://localhost:5341"

func TestLogger_INFORMATION(t *testing.T) {

	logger, _ := GetLogger(baseURL, "")

	logger.Information("Logging test message", NewProperties())

	logger.Close()

}

func TestLogger_WARNING(t *testing.T) {

	logger, _ := GetLogger(baseURL, "")

	logger.Warning("Logging test message", NewProperties())

	logger.Close()

}

func TestLogger_WithArgs(t *testing.T) {

	logger, _ := GetLogger(baseURL, "")

	var props = NewProperties()
	props.AddProperty("GUID", "11AE3484-9CD4-4332-98B1-145AAEBEACAB")
	props.AddProperty("String", "SEQ")
	props.AddProperty("Key", "Value")

	logger.Warning("Message with args", props)

	logger.Close()

}

func BenchmarkLogger_WithArgs_100times(b *testing.B) {

	logger, _ := GetLogger(baseURL, "")

	var props = NewProperties()
	props.AddProperty("GUID", "11AE3484-9CD4-4332-98B1-145AAEBEACAB")
	props.AddProperty("String", "SEQ")
	props.AddProperty("Key", "Value")
	props.AddProperty("Um", "Dois")

	for index := 0; index < b.N; index++ {
		logger.Warning(fmt.Sprintf("Message with args %d", index), props)
	}

	logger.Close()

}

func TestLogger_URLError(t *testing.T) {

	_, err := GetLogger("", "")

	if err != nil {
		t.Log("Worked")
	}

}

// TestLogger_URLError_Fail tests if even if passing a empty URL the validation fails
func TestLogger_URLError_Fail(t *testing.T) {

	_, err := GetLogger("", "")

	if err == nil {
		t.FailNow()
	}
}

func TestLogger_WithAPIKey(t *testing.T) {

	logger, _ := GetLogger(baseURL, "UWL08yUfTyw4FgXbSR")

	var props = NewProperties()
	props.AddProperty("GUID", "11AE3484-9CD4-4332-98B1-145AAEBEACAB")
	props.AddProperty("String", "SEQ")
	props.AddProperty("Key", "Value")

	logger.Warning("Message with APIKEY", props)

	logger.Close()

}

func TestLogger_DefaultProperties(t *testing.T) {

	logger, _ := GetLogger(baseURL, "")

	logger.SetDefaultProperties(map[string]interface{}{
		"Application": "TEST",
		"Teste":       "TEST",
	})

	logger.Information("WithDefaultProperties", NewProperties())

	logger.Close()
}

func TestLogger_ObjectOnProperty(t *testing.T) {
	type todo struct {
		Description string
		ID          int
	}
	type object struct {
		Name string
		Age  int
		ToDo []todo
	}

	c := make([]object, 0, 0)
	c = append(c, object{
		Age:  22,
		Name: "Munir",
		ToDo: []todo{{
			Description: "Some description",
			ID:          1,
		}, {
			Description: "Another description",
			ID:          2,
		}},
	})

	c = append(c, object{
		Age:  28,
		Name: "Moneda",
		ToDo: []todo{{
			Description: "Some description",
			ID:          1,
		}, {
			Description: "Another description",
			ID:          2,
		}},
	})

	logger, _ := GetLogger(baseURL, "")

	logger.SetDefaultProperties(map[string]interface{}{
		"Application": "TEST",
		"Teste":       c,
	})

	logger.Information("WithObjectProperties", NewProperties())

	logger.Close()
}
