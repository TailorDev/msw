package command

import (
	"flag"
	"fmt"
	"net/http"
	"strings"

	"github.com/TailorDev/msw/issue"
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

	i, err := issue.Parse(args[0])
	if err != nil {
		c.UI.Error(fmt.Sprintf("%s", err))
		return 1
	}

	errors := []string{}

	if i.WelcomeText == "" {
		errors = append(errors, "Welcome text is empty")
	}

	done := make(chan string)
	wait := 0
	nbLinks := 0
	for _, category := range i.Categories {
		if len(category.Links) == 0 {
			errors = append(errors, fmt.Sprintf("No link found for category '%s'", category.Title))
			continue
		}
		nbLinks = nbLinks + len(category.Links)

		for idx, link := range category.Links {
			if link.Name == "" {
				errors = append(errors, fmt.Sprintf(
					"No name given for link #%d in category '%s'",
					idx+1,
					category.Title,
				))
			}
			if link.URL == "" {
				errors = append(errors, fmt.Sprintf(
					"No URL given for link #%d in category '%s'",
					idx+1,
					category.Title,
				))
			} else {
				wait++
				go testLinkURL(link.URL, done)
			}
			if link.Abstract == "" {
				errors = append(errors, fmt.Sprintf(
					"No abstract given for link #%d in category '%s'",
					idx+1,
					category.Title,
				))
			}
		}
	}

	for i := 0; i < wait; i++ {
		if e := <-done; e != "" {
			errors = append(errors, e)
		}
	}

	if nbLinks > issue.MaxLinks {
		errors = append(errors, fmt.Sprintf(
			"An issue should not have more than %d links, found: %d",
			issue.MaxLinks,
			nbLinks,
		))
	}

	if len(errors) == 0 {
		c.UI.Output("Everything looks good 👍")
		return 0
	}

	for _, e := range errors {
		c.UI.Output(e)
	}

	return 1
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

func testLinkURL(url string, done chan string) {
	resp, err := http.Head(url)
	if err != nil {
		done <- fmt.Sprintf("Error while trying to perform a HEAD request: %s", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		done <- fmt.Sprintf("Could not reach URL = '%s'", url)
		return
	}

	done <- ""
}
