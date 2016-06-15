package command

import (
	"bytes"
	"flag"
	"fmt"
	"html/template"
	"strings"

	"github.com/TailorDev/msw/parser"
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
	}).Parse(issueHTML)
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
	markdown = bytes.TrimPrefix(markdown, []byte("<p>"))
	markdown = bytes.TrimSuffix(markdown, []byte("</p>\n"))

	return template.HTML(markdown)
}

const issueHTML = `<h1>Modern Science Weekly &mdash; Issue #{{ .Number }} &mdash; {{ .Date.Format "01/02/2006" }}</h1>

<p style="text-align: justify;">{{ .WelcomeText | markdown }}</p>
<p>&nbsp;</p>
{{ range $categorie := .Categories }}
<hr>
{{- range .Links }}
<h3 style="margin-top: 2rem;">{{ $categorie.Title }} // <a href="{{ .URL }}">{{ .Name }} &rarr;</a></h3>
<p style="text-align: justify;">{{ .Abstract | markdown }}</p>
{{ end -}}
{{- end -}}
<p>&nbsp;</p>
<hr>
<p style="text-align: justify;">If you received this email directly then you're already signed up, thanks! If however someone forwarded this email to you and you'd like to get it each week then you can subscribe at <a href="https://tinyletter.com/ModernScienceWeekly">https://tinyletter.com/ModernScienceWeekly</a>.</p>

<p style="text-align: center;">
    <img alt="Modern Science Weekly" class="tl-email-image" data-id="798765" height="100" src="http://gallery.tinyletterapp.com/c66e3e64ae08f8cd0d8e37710e3a5aef07eb6730/images/82443a39-2712-410f-ad7d-632b7fe2f63d.jpg" style="width: 100px; max-width: 100px;" width="100">
</p>`
