package simplejson

import (
	"fmt"
	"testing"
)

func TestSimpleJsonCmd(t *testing.T) {
	t.Run("ValidJSON", func(t *testing.T) {
		err := RunSimeleJson()
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
	})

	t.Run("InvalidJSON - wrong value type", func(t *testing.T) {
		jsonData := `
		{
			"field1": "Hello",
			"field2": "invalid"
		}`

		err := RunSimeleJsonWithJSONData(jsonData)
		if err == nil {
			t.Error("Expected error, got nil")
		} else {
			fmt.Printf("Error: %v\n", err)
		}
	})

	t.Run("InvalidJSON - value NaN", func(t *testing.T) {
		jsonData := `
		{
			"field1": "Hello",
			"field2": NaN
		}`

		err := RunSimeleJsonWithJSONData(jsonData)
		if err == nil {
			t.Error("Expected error, got nil")
		} else {
			fmt.Printf("Error: %v\n", err)
		}
	})


	t.Run("InvalidJSON - number out of range", func(t *testing.T) {
		jsonData := `
		{
			"field1": "Hello",
			"field2": 41
		}`

		err := RunSimeleJsonWithJSONData(jsonData)
		if err == nil {
			t.Error("Expected error, got nil")
		} else {
			fmt.Printf("Error: %v\n", err)
		}
	})
}
