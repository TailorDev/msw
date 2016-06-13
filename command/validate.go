package command

import (
	"flag"
	"fmt"
	"strings"

	"github.com/TailorDev/msw/parser"
	"github.com/mitchellh/cli"
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

	issue, err := parser.Parse(args[0])
	if err != nil {
		c.Ui.Error(fmt.Sprintf("%s", err))
		return 1
	}

	for _, categorie := range issue.Categories {
		c.Ui.Output(fmt.Sprintf("%v", categorie.Title))
	}

	//resp, err := http.Get("")

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
