package command

import (
	"strings"

	"github.com/mitchellh/cli"
)

// NewCommand is a Command that creates a new empty YAML file to prepare a new
// issue.
type NewCommand struct {
	UI cli.Ui
}

// Run runs the code of the comand.
func (c *NewCommand) Run(args []string) int {
	return 0
}

// Help returns the description of the command.
func (*NewCommand) Help() string {
	helpText := `
Usage: msw new [options] ISSUE_NUMBER

  This command creates a new empty YAML file to prepare a new issue.

Options:

  -date=<date>			The date of the issue.

  -directory=path		The directory where to write the generated file.

`
	return strings.TrimSpace(helpText)
}

// Synopsis returns the short description of the command.
func (*NewCommand) Synopsis() string {
	return "create a new empty YAML file to prepare a new issue."
}
