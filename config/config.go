package config

import (
	"io/ioutil"
	"path/filepath"

	"github.com/medhoover/gom/logger"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

// Defining the Configuration type
type ConfigInstance struct {
	Name     string                 `yaml:"name,omitempty" validate:"string,required"`
	Commands map[string]CommandType `yaml:"commands,flow,omitempty"`
}

type parser interface {
	new(path string) (*ConfigInstance, error)
}

// Read, parse and validate the config file
func (ci *ConfigInstance) new(path string) (*ConfigInstance, error) {

	// Read config file
	data, err := ioutil.ReadFile(path)
	if err != nil {
		absPath, _ := filepath.Abs(path)
		return nil, errors.Errorf(
			"Unable to read file %s\nCreate a configuration file first",
			absPath,
		)
	}

	// Parse the file to the ConfigInstance value
	err = yaml.Unmarshal([]byte(data), &ci)
	if err != nil {
		return nil, errors.Wrap(err, "Invalid file structure")
	}

	// Validate the ConfigInstance structure
	if valid, err := ci.Validate(); valid == false {
		return nil, errors.Wrap(err, "Invalid file structure")
	}
	return ci, nil
}

// Export a configuration instance
func New(path string) *ConfigInstance {
	// create a new configuration instance
	var ci *ConfigInstance
	ci, err := ci.new(path)
	if err != nil {
		logger.Error(err)
	}
	return ci
}
