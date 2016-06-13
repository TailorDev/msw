package parser

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"regexp"
	"time"

	"gopkg.in/yaml.v2"
)

// Parse allows to parse a YAML file and returns a map
func Parse(filename string) (Issue, error) {
	issue := Issue{}

	f, _ := filepath.Abs(filename)
	yamlFile, err := ioutil.ReadFile(f)
	if err != nil {
		return issue, err
	}

	r, _ := regexp.Compile("^.*([0-9]{4}-[0-9]{2}-[0-9]{2})\\.yml$")
	matches := r.FindStringSubmatch(f)
	if len(matches) != 2 {
		return issue, fmt.Errorf("Invalid filename (%s), should match 'YYYY-MM-DD.yml'", filename)
	}

	date, err := time.Parse("2006-01-02", matches[1])
	if err != nil {
		return issue, err
	}

	if err := yaml.Unmarshal(yamlFile, &issue); err != nil {
		return issue, err
	}

	issue.Date = date

	return issue, nil
}
