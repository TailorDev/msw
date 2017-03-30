package command

import (
	"flag"
	"fmt"
	"strings"

	"github.com/TailorDev/msw/buffer"
	"github.com/TailorDev/msw/config"
	"github.com/TailorDev/msw/issue"
	"github.com/mitchellh/cli"
)

// BufferPushCommand is a Command that pushes links to Buffer.com's queue.
type BufferPushCommand struct {
	UI cli.Ui
}

// Run runs the code of the comand.
func (c *BufferPushCommand) Run(args []string) int {
	var apply bool
	cmdFlags := flag.NewFlagSet("buffer push", flag.ContinueOnError)
	cmdFlags.Usage = func() { c.UI.Output(c.Help()) }
	cmdFlags.BoolVar(&apply, "apply", false, "")
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

	conf, err := config.LoadDefaultConfig()
	if err != nil {
		c.UI.Error(fmt.Sprintf("%s", err))
		return 1
	}

	if conf.Buffer.AccessToken == "" {
		c.UI.Error("You must specify an access token in the configuration file.")
		return 1
	}

	if len(conf.Buffer.ProfileIDs) == 0 {
		c.UI.Error("You must specify at least one profile ID in the configuration file.")
		return 1
	}

	if !apply {
		c.UI.Output("Re-run the command with `-apply` to actually push to Buffer.com\n")
	}

	cli := buffer.NewClient(conf.Buffer.AccessToken)
	for _, category := range i.Categories {
		for _, link := range category.Links {
			if link.Name != "" && link.URL != "" {
				text := link.GetBufferText()

				if apply {
					updates, err := cli.Push(text, conf.Buffer.ProfileIDs)
					if err != nil {
						c.UI.Error(fmt.Sprintf("%s", err))
						return 1
					}

					for _, u := range updates {
						c.UI.Output(fmt.Sprintf("[%s] %s", u.ProfileService, u.Text))
					}
				} else {
					c.UI.Output(fmt.Sprintf("[?] %s", text))
				}
			}
		}
	}

	return 0
}

// Help returns the description of the command.
func (*BufferPushCommand) Help() string {
	helpText := `
Usage: msw buffer push [options] FILENAME

  This command pushes each entry of an issue to Buffer.com's queue. You need
  a configuration file with Buffer credentials ('~/.msw/msw.yml').

Options:

  -apply				Push to Buffer.com's queue.

`
	return strings.TrimSpace(helpText)
}

// Synopsis returns the short description of the command.
func (*BufferPushCommand) Synopsis() string {
	return "push links to Buffer.com's queue"
}
