package config

import (
	"fmt"
	"os"
	"os/exec"
	"reflect"
	"runtime"
	"strings"

	"github.com/medhoover/gom/logger"
)

type anyType interface{}

type Command interface {
	Execute([]string, int) error
	populate(interface{}) (*Command, error)
}

type CommandType struct {
	commands    []string
	subCommands map[string]*CommandType
}

// gom talks with the system shell to support POSIX scripting
var (
	sysShell      = "sh"
	sysCommandArg = "-c"
)

func init() {
	if runtime.GOOS == "windows" {
		sysShell = "cmd"
		sysCommandArg = "/c"
	}
}

// // For proper display
// func (c CommandType) String() string {
// 	if c.commands != nil {
// 		return fmt.Sprintf("%v ", c.commands)
// 	} else {
// 		return fmt.Sprintf("%v ", c.subCommands)
// 	}
// }

// Execute a command, it accepts a path slice and the related command index
func (c *CommandType) Execute(path []string, index int) error {
	// increment index to point to the current command in the path slice
	index++
	// Execute commands if they exist. No need to check for the subCommands map
	if len(c.commands) > 0 {
		for _, command := range c.commands {
			// Build command arguments
			args := command + " " + strings.Join(path[index:], " ")
			// Talk to system shell. Example (Unix): sh -c args
			cmd := exec.Command(sysShell, sysCommandArg, args)
			cmd.Stdin = os.Stdout
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			logger.Command(args)
			if err := cmd.Start(); err != nil {
				return err
			}
			if err := cmd.Wait(); err != nil {
				return err
			}
		}
	} else {
		// Before checking subCommands property, make sure it still has path left
		if len(path) > index {
			// Execute subCommand if found
			if subCommand, exist := c.subCommands[path[index]]; exist {
				return subCommand.Execute(path, index)
			}
			return fmt.Errorf("Command is not defined")
		}
		return fmt.Errorf("Please specify a sub-command")
	}
	return nil
}

// Custom Unmarshal method for CommandType
// Populate the type structure based on the "commands" property from the YAML file
func (c *CommandType) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var m anyType
	err := unmarshal(&m)
	if err != nil {
		return err
	}
	return c.populate(m)
}

// Recursive method to populate a Command value
func (c *CommandType) populate(v interface{}) error {
	// Populate command struct based on the value type
	switch reflect.TypeOf(v).Kind() {
	// Append the string to the commands property
	case reflect.String:
		c.commands = []string{reflect.ValueOf(v).Interface().(string)}
	case reflect.Slice:
		// Extract commands (string typed) and append it to commands property
		slice := reflect.ValueOf(v)
		c.commands = make([]string, slice.Len())
		for i := 0; i < slice.Len(); i++ {
			str, ok := slice.Index(i).Interface().(string)
			if !ok {
				return fmt.Errorf("\nInvalid command %v. Arrays should contain only string elements",
					slice.Index(i))
			}
			c.commands[i] = str
		}
	case reflect.Map:
		m := reflect.ValueOf(v).Interface().(map[interface{}]interface{})
		c.subCommands = make(map[string]*CommandType)
		for key, value := range m {
			if value == nil {
				return fmt.Errorf("\nCommand '%s' cannot be empty", key)
			}
			c.subCommands[key.(string)] = &CommandType{}
			if err := c.subCommands[key.(string)].populate(value); err != nil {
				return err
			}
		}
	default:
		return fmt.Errorf("\nInvalid command value %v. Type %v is not supported in this command",
			v,
			reflect.TypeOf(v).Kind())
	}
	return nil
}
