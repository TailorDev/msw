package command

import (
	"flag"
	"fmt"
	"strings"

	"github.com/TailorDev/msw/parser"
	"github.com/mitchellh/cli"
)

// ValidateCommand is a Command that validates a YAML file.
type ValidateCommand struct {
	UI cli.Ui
}

// Run runs the code of the comand.
func (c *ValidateCommand) Run(args []string) int {
	cmdFlags := flag.NewFlagSet("validate", flag.ContinueOnError)
	cmdFlags.Usage = func() { c.UI.Output(c.Help()) }
	if err := cmdFlags.Parse(args); err != nil {
		return 1
	}

	args = cmdFlags.Args()
	if len(args) != 1 {
		cmdFlags.Usage()
		return 1
	}

	issue, err := parser.Parse(args[0])
	if err != nil {
		c.UI.Error(fmt.Sprintf("%s", err))
		return 1
	}

	for _, categorie := range issue.Categories {
		c.UI.Output(fmt.Sprintf("%v", categorie.Title))
	}

	//resp, err := http.Get("")

	return 0
}

// Help returns the description of the command.
func (*ValidateCommand) Help() string {
	helpText := `
Usage: msw validate FILENAME

  This command checks whether a filename contains a valid issue.

`
	return strings.TrimSpace(helpText)
}

// Synopsis returns the short description of the command.
func (*ValidateCommand) Synopsis() string {
	return "check that an issue is valid"
}
