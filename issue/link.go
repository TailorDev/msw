package issue

import "strings"

var (
	// DefaultBufferTemplate is the default Buffer template used to construct tweets.
	DefaultBufferTemplate = "%name%: %url%."
)

// Link is a structure representing a link entry in a newsletter issue.
type Link struct {
	Name     string
	URL      string `yaml:"url"`
	Abstract string
	Buffer   string
}

func (l *Link) GetBufferText() string {
	t := l.Buffer
	if t == "" {
		t = DefaultBufferTemplate
	}

	name := strings.Replace(l.Name, "\n", " ", -1)
	name = strings.TrimSpace(name)

	url := strings.TrimSpace(l.URL)

	t = strings.Replace(t, "%name%", name, -1)
	t = strings.Replace(t, "%url%", url, -1)

	return strings.TrimSpace(t)
}
