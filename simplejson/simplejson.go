package simplejson

import (
	"encoding/json"
	"fmt"
)

type MyStruct struct {
	Field1 string `json:"field1"`
	Field2 int    `json:"field2"`
}

func RunSimeleJson() error {
	jsonData := `{"field1": "Hello", "field2": 42}`

	var myInstance MyStruct

	err := json.Unmarshal([]byte(jsonData), &myInstance)
	if err != nil {
		return fmt.Errorf("error unmarshalling JSON: %v", err)
	}

	// Print the values of the struct fields
	fmt.Printf("Field1: %s\n", myInstance.Field1)
	fmt.Printf("Field2: %d\n", myInstance.Field2)

	return nil
}
