package command

import (
	"flag"
	"fmt"
	"html/template"
	"os"
	"strings"

	"github.com/TailorDev/msw/parser"
	"github.com/TailorDev/msw/tpl"
	"github.com/mitchellh/cli"
	"github.com/russross/blackfriday"
)

type GenerateCommand struct {
	Ui cli.Ui
}

func (c *GenerateCommand) Run(args []string) int {
	cmdFlags := flag.NewFlagSet("generate", flag.ContinueOnError)
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

	t, err := template.New("issue").Funcs(template.FuncMap{
		"markdown": func(s string) template.HTML {
			return template.HTML(blackfriday.MarkdownBasic([]byte(s)))
		},
	}).Parse(tpl.IssueHTML)
	if err != nil {
		c.Ui.Error(fmt.Sprintf("Error parsing template: %s", err))
		return 1
	}

	if err = t.Execute(os.Stdout, issue); err != nil {
		c.Ui.Error(fmt.Sprintf("Error generating HTML: %s", err))
		return 1
	}

	return 0
}

func (*GenerateCommand) Help() string {
	helpText := `
Usage: msw generate FILENAME

  This command generates HTML for Tinyletter from a YAML file.

`
	return strings.TrimSpace(helpText)
}

func (*GenerateCommand) Synopsis() string {
	return "generate HTML for Tinyletter from a YAML file"
}
