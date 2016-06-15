package issue

import "time"

// MaxLinks represents the maximal number of links per issue.
const MaxLinks int = 8

// Issue is a structure representing a newsletter issue.
type Issue struct {
	Number      int
	WelcomeText string `yaml:"welcome_text"`
	Date        time.Time
	Categories  []struct {
		Title string
		Links []struct {
			Name     string
			URL      string `yaml:"url"`
			Abstract string
		}
	}
}
