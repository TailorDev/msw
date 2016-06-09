package command

import (
	"flag"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/mitchellh/cli"
	"gopkg.in/yaml.v2"
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

	filename, _ := filepath.Abs(args[0])
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		c.Ui.Error(fmt.Sprintf("Error reading file: %s", err))
		return 1
	}

	issue := map[string]interface{}{}
	if err := yaml.Unmarshal(yamlFile, &issue); err != nil {
		c.Ui.Error(fmt.Sprintf("Error parsing file: %s", err))
		return 1
	}

	t, err := template.New("issue.html").ParseFiles("template/issue.html")
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
