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

type builder interface {
	populate(interface{}) (*command, error)
}

type command struct {
	lines       []string
	subCommands map[string]*command
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

// Execute a command, it accepts a path slice and the related command index
func (c *command) Execute(args []string) error {
	// Execute commands if they exist. No need to check for the subCommands map
	if len(c.lines) > 0 {
		for _, command := range c.lines {
			// If user passes arguments to command, join it.
			command = command + " " + strings.Join(args, " ")
			// Talk to system shell. Example (Unix): sh -c args
			cmd := exec.Command(sysShell, sysCommandArg, command)
			cmd.Stdin = os.Stdout
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			logger.Command(command)
			if err := cmd.Start(); err != nil {
				return err
			}
			if err := cmd.Wait(); err != nil {
				return err
			}
		}
	} else {
		// Before checking subCommands property, make sure it still has args left
		if len(args) > 0 {
			// Execute subCommand if found
			if subCommand, exist := c.subCommands[args[0]]; exist {
				return subCommand.Execute(args[1:])
			}
			return fmt.Errorf(("Command '%s' is not defined"), args[0])
		}
		return fmt.Errorf("Please specify a sub-command")
	}
	return nil
}

// UnmarshalYAML custom Unmarshal method for Command
// Populate the type structure based on the "commands" property from the YAML file
func (c *command) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var m anyType
	err := unmarshal(&m)
	if err != nil {
		return err
	}
	return c.populate(m)
}

// Recursive method to populate a Command value
func (c *command) populate(v interface{}) error {
	// Populate command struct based on the value type
	switch reflect.TypeOf(v).Kind() {
	// Append the string to the commands property
	case reflect.String:
		c.lines = []string{reflect.ValueOf(v).Interface().(string)}
	case reflect.Slice:
		// Extract commands (string typed) and append it to commands property
		slice := reflect.ValueOf(v)
		c.lines = make([]string, slice.Len())
		for i := 0; i < slice.Len(); i++ {
			str, ok := slice.Index(i).Interface().(string)
			if !ok {
				return fmt.Errorf("\nInvalid command %v. Arrays should contain only string elements",
					slice.Index(i))
			}
			c.lines[i] = str
		}
	case reflect.Map:
		m := reflect.ValueOf(v).Interface().(map[interface{}]interface{})
		c.subCommands = make(map[string]*command)
		for key, value := range m {
			if value == nil {
				return fmt.Errorf("\nCommand '%s' cannot be empty", key)
			}
			c.subCommands[key.(string)] = &command{}
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
