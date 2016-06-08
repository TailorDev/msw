package command

import (
	"flag"
	"github.com/mitchellh/cli"
	"strings"
)

type ValidateCommand struct {
	Ui cli.Ui
}

func (c *ValidateCommand) Run(args []string) int {
	cmdFlags := flag.NewFlagSet("validate", flag.ContinueOnError)
	cmdFlags.Usage = func() { c.Ui.Output(c.Help()) }
	if err := cmdFlags.Parse(args); err != nil {
		return 1
	}

	args = cmdFlags.Args()
	if len(args) != 1 {
		cmdFlags.Usage()
		return 1
	}

	return 0
}

func (*ValidateCommand) Help() string {
	helpText := `
Usage: msw validate FILENAME

  This command checks whether a filename contains a valid issue.

`
	return strings.TrimSpace(helpText)
}

func (*ValidateCommand) Synopsis() string {
	return "check that an issue is valid"
}
