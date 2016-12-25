package cli

import (
	"flag"
	"fmt"

	"github.com/fatih/color"
)

var (
	FlagCount   = 0
	ArgCount    = 0
	FilePath    = "./gom.yaml"
	VersionFlag = false
)

func init() {

	// Define how help is displayed
	flag.Usage = func() {
		yellow := color.New(color.FgHiYellow).SprintFunc()
		red := color.New(color.FgRed).SprintFunc()
		fmt.Print(
			"-----------\n",
			yellow("gom 0.1.0\n"),
			"Usage:\tgom [Options] Commands... [Args_Options...]\n\n",
			red("Options:\n"),
			"\t-v\tShow gom version.\n",
			"\t-f\tSet configuration file.\tExample: -f=conf.yaml\n",
			red("Commands:\n"),
			"\tKey path as specified in the commands property of the YAML config file.\n",
			red("Args_Options:\n"),
			"\tPasse your arguments to the specified command\n",
			"Example:\n",
			"\tConfig File:\n",
			"\t-----\n",
			"\tname: projectName\n",
			"\tcommands:\n",
			"\t  list:\n",
			"\t    names: ls\n",
			"\t    all: ls -al\n",
			"\t  greet:\n",
			"\t    - echo Hello!\n",
			"\t    - echo Hallo!\n",
			"\t    - echo Bonjour!\n",
			"\t-----\n",
			"\t# Define nested commands\n",
			"\t> gom list names => File1 Folder1\n",
			"\t# Define sequence commands\n",
			"\t> gom greet =>      Hello!\n",
			"\t                    Hallo!\n",
			"\t                    Bonjour!\n",
			"\t# Passe optional arguments\n",
			"\t> gom names -f =>   File1 .. .\n\n",
			"-----------\n",
			"Author:\tMohamed GHARSALLAH https://gharsallah.com\n",
			"For more details check: https://github.com/medhoover/gom\n",
		)
	}

	// Set flags
	flag.BoolVar(&VersionFlag, "v", false, "Show gom version")
	flag.StringVar(&FilePath, "f", FilePath, "Set configuration file path")

	// Parse the defined flags
	flag.Parse()
	FlagCount = flag.NFlag()
	ArgCount = flag.NArg()
}
