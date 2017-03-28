package issue

import "time"

// MaxLinks represents the maximal number of links per issue.
const MaxLinks int = 8

// DefaultCategories contains the titles of the default categories
var DefaultCategories = []string{
	"Open Science & Data",
	"Tools for Scientists",
	"Cutting-edge Science",
	"Beyond Academia",
}

// Issue is a structure representing a newsletter issue.
type Issue struct {
	Number      int
	WelcomeText string `yaml:"welcome_text"`
	Date        time.Time
	Categories  []struct {
		Title string
		Links []Link
	}
}
