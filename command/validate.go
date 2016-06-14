package command

import (
	"flag"
	"fmt"
	"net/http"
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

	if issue.WelcomeText == "" {
		c.UI.Output("Welcome text is empty")
	}

	for _, category := range issue.Categories {
		if len(category.Links) == 0 {
			c.UI.Output(fmt.Sprintf("No link found for category '%s'", category.Title))
			continue
		}

		done := make(chan bool)
		wait := 0
		for idx, link := range category.Links {
			if link.Name == "" {
				c.UI.Output(fmt.Sprintf(
					"No name given for link #%d in category '%s'",
					idx+1,
					category.Title,
				))
			}
			if link.URL == "" {
				c.UI.Output(fmt.Sprintf(
					"No URL given for link #%d in category '%s'",
					idx+1,
					category.Title,
				))
			} else {
				wait++
				go testLinkURL(link.URL, c.UI, done)
			}
		}
		for i := 0; i < wait; i++ {
			<-done
		}
	}

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

func testLinkURL(url string, ui cli.Ui, done chan bool) {
	resp, err := http.Head(url)
	if err != nil {
		ui.Error(fmt.Sprintf("Error while trying to perform a HEAD request: %s", err))
	} else {
		defer resp.Body.Close()

		if resp.StatusCode != 200 {
			ui.Output(fmt.Sprintf("Could not reach URL = '%s'", url))
		}
	}

	done <- true
}
