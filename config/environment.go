package config

import (
	"os"
	"reflect"

	"github.com/pkg/errors"
)

type Environment struct {
	variables map[string]string
}

var invalidEnvError = errors.New("\nEnvironment definition should have the following syntax:\nenv:\n  envName:\n    valName: value (string)\n    ...")

// Set the variables of the map receiver as environment variables
func (env *Environment) Set() error {
	for key, value := range env.variables {
		if err := os.Setenv(key, value); err != nil {
			return err
		}
	}
	return nil
}

// Custom Unmarshal method for Environment
func (env *Environment) UnmarshalYAML(unmarshal func(interface{}) error) error {
	// Prepare a variable to cast any type
	var e anyType

	// Parse 'env' value to the variable
	unmarshal(&e)

	// Check if the variable type is valid
	if e == nil || reflect.TypeOf(e).Kind() != reflect.Map {
		return errors.Wrap(invalidEnvError, "An empty or invalid environment")
	}

	// Prepare 'variables' property
	env.variables = make(map[string]string)

	// Iterate over the map fields
	for key, value := range reflect.ValueOf(e).Interface().(map[interface{}]interface{}) {
		// The value must be of type string
		if value == nil || !reflect.TypeOf(value).AssignableTo(reflect.TypeOf("")) {
			return errors.Wrapf(invalidEnvError, "Value of environment property '%v' should be a string", key)
		}
		// Add the value to the map
		env.variables[key.(string)] = value.(string)
	}
	return nil
}
