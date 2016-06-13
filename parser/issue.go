package parser

import "time"

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
