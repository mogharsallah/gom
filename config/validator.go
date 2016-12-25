package config

import (
	"fmt"
	"reflect"
	"strings"
)

// Name of the struct tag used in examples.
const tagName = "validate"

// Generic data validator.
type Validator interface {
	// Validate method performs validation and returns result and optional error.
	Validate(interface{}) (bool, error)
}

// DefaultValidator does not perform any validations.
type DefaultValidator struct {
}

func (v DefaultValidator) Validate(val interface{}) (bool, error) {
	return true, nil
}

// Returns validator struct corresponding to validation type
func getValidatorFromTag(tag string) Validator {
	args := strings.Split(tag, ",")
	argsMap := make(map[string]bool)
	for _, value := range args {
		argsMap[value] = true
	}
	switch args[0] {
	case "string":
		validator := StringValidator{
			required: argsMap["required"] == true,
		}
		return validator
	}

	return DefaultValidator{}
}

// Performs actual data validation using validator definitions on the struct
func (ci ConfigInstance) Validate() (bool, error) {

	// ValueOf returns a Value representing the run-time data
	v := reflect.ValueOf(ci)
	for i := 0; i < v.NumField(); i++ {
		// Get the field tag value
		tag := v.Type().Field(i).Tag.Get(tagName)

		// Skip if tag is not defined or ignored
		if tag == "" || tag == "-" {
			continue
		}

		// Get a validator that corresponds to a tag
		validator := getValidatorFromTag(tag)

		// Perform validation
		valid, err := validator.Validate(v.Field(i).Interface())

		// Append error to results
		if !valid && err != nil {
			return false, fmt.Errorf("%s %s", v.Type().Field(i).Name, err.Error())
		}
	}
	return true, nil
}

// Custom validators

// StringValidator validates string presence and/or its length.
type StringValidator struct {
	required bool
}

func (v StringValidator) Validate(val interface{}) (bool, error) {

	l := len(val.(string))

	if l == 0 && v.required == true {
		return false, fmt.Errorf("cannot be blank")
	}

	return true, nil
}
