package cli

import (
	"os"
	"strings"

	"github.com/medhoover/gom/config"
	"github.com/medhoover/gom/logger"
	"github.com/pkg/errors"
)

type Action struct {
	Path  []string
	Index int
}

func New() {

	// Show gom version if -v flag was present
	if VersionFlag {
		logger.Info("Version 0.1.0\nPlease check https://github.com/medhoover/gom for new updates")
		os.Exit(1)
	}

	// Show help if no arguments were entered
	if ArgCount == 0 {
		Usage()
		os.Exit(1)
	}

	a := Action{Path: os.Args[FlagCount+1:]}
	ci := config.New(FilePath)
	launche(a, ci)
}

func launche(a Action, ci *config.ConfigInstance) {

	if command, exist := ci.Commands[a.Path[a.Index]]; exist {
		if err := command.Execute(a.Path, a.Index); err != nil {
			logger.Error(
				errors.Wrapf(
					err,
					"Invalid command '%s'",
					strings.Join(a.Path[a.Index:], " "),
				),
			)
		}
	} else {
		logger.Error(errors.Errorf("Command '%s' is not defined", a.Path[a.Index]))
	}
}
