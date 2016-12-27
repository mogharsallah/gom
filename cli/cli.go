package cli

import (
	"os"
	"strings"

	"github.com/medhoover/gom/config"
	"github.com/medhoover/gom/logger"
	"github.com/pkg/errors"
)

func New() {

	// Show gom version if -v flag was present
	if VersionFlag {
		logger.Info("Version 0.1.0\nPlease check https://github.com/medhoover/gom for new updates")
		os.Exit(1)
	}
	// Show help if no arguments were entered
	if len(os.Args) < FlagCount+1 {
		Usage()
		os.Exit(1)
	}

	path := os.Args[FlagCount+1:]
	index := 0
	ci := config.New(FilePath)
	launche(path, index, ci)
}

func launche(path []string, index int, ci *config.ConfigInstance) {

	if command, exist := ci.Commands[path[index]]; exist {
		if err := command.Execute(path, index); err != nil {
			logger.Error(
				errors.Wrapf(
					err,
					"Command '%s' Failed",
					strings.Join(path[index:], " "),
				),
			)
		}
	} else {
		logger.Error(errors.Errorf("Command '%s' Failed: Command is not defined", path[index]))
	}
}
