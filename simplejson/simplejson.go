package simplejson

import (
	"encoding/json"
	"fmt"

	"github.com/go-playground/validator/v10"
)

type MyStruct struct {
	Field1 string `json:"field1" validate:"required"`
	Field2 int    `json:"field2" validate:"gte=0"`
}

func RunSimeleJson() error {
	jsonData := `
	{
		"field1": "Hello",
		"field2": 42
	}`

	return RunSimeleJsonWithJSONData(jsonData)
}

func RunSimeleJsonWithJSONData(jsonData string) error {
	var myInstance MyStruct

	err := json.Unmarshal([]byte(jsonData), &myInstance)
	if err != nil {
		return fmt.Errorf("error unmarshalling JSON: %v", err)
	}

	if err := validateStruct(myInstance); err != nil {
		return fmt.Errorf("error validating JSON: %v", err)
	}

	fmt.Printf("Field1: %s\n", myInstance.Field1)
	fmt.Printf("Field2: %d\n", myInstance.Field2)

	return nil
}

func validateStruct(s MyStruct) error {
	validate := validator.New()

	if err := validate.Struct(s); err != nil {
		return fmt.Errorf("validation error: %v", err)
	}

	return nil
}
