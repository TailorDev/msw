package command

import (
	"flag"
	"fmt"
	"strings"

	"github.com/TailorDev/msw/buffer"
	"github.com/TailorDev/msw/config"
	"github.com/mitchellh/cli"
)

// BufferInfoCommand is a Command that outputs Buffer.com information.
type BufferInfoCommand struct {
	UI   cli.Ui
	Conf config.Config
}

// Run runs the code of the comand.
func (c *BufferInfoCommand) Run(args []string) int {
	cmdFlags := flag.NewFlagSet("buffer-info", flag.ContinueOnError)
	cmdFlags.Usage = func() { c.UI.Output(c.Help()) }
	if err := cmdFlags.Parse(args); err != nil {
		return 1
	}

	if c.Conf.Buffer.AccessToken == "" {
		c.UI.Error("You must specify an access token in the configuration file.")
		return 1
	}

	if len(c.Conf.Buffer.ProfileIDs) == 0 {
		c.UI.Error("You must specify at least one profile ID in the configuration file.")
		return 1
	}

	cli := buffer.NewClient(c.Conf.Buffer.AccessToken)
	for _, id := range c.Conf.Buffer.ProfileIDs {
		p, err := cli.GetProfile(id)
		if err != nil {
			c.UI.Error(fmt.Sprintf("%s", err))
			continue
		}

		c.UI.Output(fmt.Sprintf("\nProfile: %s (%s)", p.FormattedUsername, p.Service))

		updates, _ := cli.GetPendingUpdates(id)

		c.UI.Output(fmt.Sprintf("Buffer : %d / %d\n", len(updates), c.Conf.Buffer.BufferSize))
		c.UI.Output("Scheduled updates:\n")

		for _, u := range updates {
			c.UI.Output(fmt.Sprintf("%-25s: %s", u.Day, u.Text))
		}
	}

	return 0
}

// Help returns the description of the command.
func (*BufferInfoCommand) Help() string {
	helpText := `
Usage: msw buffer info

  This command shows information about Buffer.com. You need a configuration
  file with Buffer credentials ('~/.msw/msw.yml').

`
	return strings.TrimSpace(helpText)
}

// Synopsis returns the short description of the command.
func (*BufferInfoCommand) Synopsis() string {
	return "output Buffer.com information"
}
