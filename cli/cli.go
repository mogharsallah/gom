/*
cli package
Handels input and initiate commands execution
*/
package cli

import (
	"os"

	flags "github.com/jessevdk/go-flags"
	"github.com/medhoover/gom/config"
	"github.com/medhoover/gom/logger"
)

const gomVersion = "0.2.0"
const filePath = "gom.yaml"

var options struct {
	Version  func() `long:"version" description:"Show gom version"`
	FilePath string `short:"f" long:"file" description:"Configuration file path" value-name:"path/to/file"`
	Env      string `short:"e" long:"environment" description:"Set environment variables from 'env' property" value-name:"Env"`
	Usage    string
}

// New reads the input arguments and execute the related command
func New() {
	// Set callback for --version flag
	options.Version = showVersion

	// Set default FilePath. If user uses the file flag, the value will change
	options.FilePath = filePath

	// Create a flag parser
	parser := flags.NewParser(&options, flags.HelpFlag|flags.PassAfterNonOption|flags.PrintErrors)

	// Define how usage section is shown inside help
	parser.Usage = "[options] command [command_options...] "

	// Parse flags and retrieve arguments
	args, err := parser.Parse()
	if err != nil {
		os.Exit(0)
	}

	// Create a new configuration instance from the file path
	ci := config.New(options.FilePath)

	// Use the given environment
	if options.Env != "" {
		ci.Set(options.Env)
	}

	// Exit if no arguments were entered
	if len(args) == 0 {
		parser.WriteHelp(os.Stdout)
		os.Exit(0)
	}

	// Execute the passed command alias
	ci.Execute(args)
}

// showVersion prints gom version, used when --version flag is issued
func showVersion() {
	logger.Info("Version " + gomVersion + "\nPlease check https://github.com/medhoover/gom for new updates")
	os.Exit(0)
}
