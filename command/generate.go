package command

import (
	"bytes"
	"flag"
	"fmt"
	"html/template"
	"strings"

	"github.com/TailorDev/msw/parser"
	"github.com/TailorDev/msw/tpl"
	"github.com/mitchellh/cli"
	"github.com/russross/blackfriday"
)

// GenerateCommand is a Command that generates HTML from YAML files.
type GenerateCommand struct {
	UI cli.Ui
}

// Run runs the code of the comand.
func (c *GenerateCommand) Run(args []string) int {
	cmdFlags := flag.NewFlagSet("generate", flag.ContinueOnError)
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

	t, err := template.New("issue").Funcs(template.FuncMap{
		"markdown": markdown,
	}).Parse(tpl.IssueHTML)
	if err != nil {
		c.UI.Error(fmt.Sprintf("Error parsing template: %s", err))
		return 1
	}

	var out bytes.Buffer
	if err = t.Execute(&out, issue); err != nil {
		c.UI.Error(fmt.Sprintf("Error generating HTML: %s", err))
		return 1
	}

	c.UI.Output(out.String())

	return 0
}

// Help returns the description of the command.
func (*GenerateCommand) Help() string {
	helpText := `
Usage: msw generate FILENAME

  This command generates HTML for Tinyletter from a YAML file.

`
	return strings.TrimSpace(helpText)
}

// Synopsis returns the short description of the command.
func (*GenerateCommand) Synopsis() string {
	return "generate HTML for Tinyletter from a YAML file"
}

func markdown(s string) template.HTML {
	markdown := blackfriday.MarkdownBasic([]byte(s))
	// remove enclosing <p> tag
	markdown = bytes.TrimLeft(markdown, "<p>")
	markdown = bytes.TrimRight(markdown, "</p>\n")

	return template.HTML(markdown)
}
