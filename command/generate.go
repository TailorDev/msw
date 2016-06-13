package command

import (
	"flag"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

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

	r, _ := regexp.Compile("^.*([0-9]{4}-[0-9]{2}-[0-9]{2})\\.yml$")
	matches := r.FindStringSubmatch(filename)
	if len(matches) != 2 {
		c.Ui.Error(fmt.Sprintf("Invalid filename (%s), should match 'YYYY-MM-DD.yml'", filename))
		return 1
	}

	date, err := time.Parse("2006-01-02", matches[1])
	if err != nil {
		c.Ui.Error(fmt.Sprintf("Error parsing date: %s", err))
		return 1
	}

	issue := map[string]interface{}{}
	if err := yaml.Unmarshal(yamlFile, &issue); err != nil {
		c.Ui.Error(fmt.Sprintf("Error parsing file: %s", err))
		return 1
	}

	issue["date"] = date

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
