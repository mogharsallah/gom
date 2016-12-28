package cli

import (
	"os"

	flags "github.com/jessevdk/go-flags"
	"github.com/medhoover/gom/config"
	"github.com/medhoover/gom/logger"
)

const GomVersion = "0.2.0"
const FilePath = "./gom.yaml"

var Options struct {
	Version  func() `long:"version" description:"Show gom version"`
	FilePath string `short:"f" long:"file" description:"Configuration file path" value-name:"FILE"`
	Usage    string
}

func New() {
	// Set callback for --version flag
	Options.Version = showVersion
	// Define how usage section is shown inside help
	Options.Usage = "[options] command [command_options...] "
	// Set default FilePath. If user uses the file flag, the value will change
	Options.FilePath = FilePath

	// Create a flag parser
	parser := flags.NewParser(&Options, flags.HelpFlag|flags.PassAfterNonOption|flags.PrintErrors)

	// Parse flags and retrieve arguments
	args, err := parser.Parse()
	if err != nil {
		os.Exit(0)
	}

	// Print help if no arguments were entered
	if len(args) == 0 {
		parser.WriteHelp(os.Stdout)
		os.Exit(0)
	}

	// Create a new configuration instance from the file path
	ci := config.New(Options.FilePath)

	// Execute the passed command alias
	ci.Execute(args)
}

// showVersion prints gom version, used when --version flag is issued
func showVersion() {
	logger.Info("Version " + GomVersion + "\nPlease check https://github.com/medhoover/gom for new updates")
	os.Exit(0)
}
