package command_test

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"testing"
	"time"

	"github.com/TailorDev/msw/command"
	"github.com/mitchellh/cli"
)

func TestNewInvalidArgs(t *testing.T) {
	ui := new(cli.MockUi)
	cmd := &command.NewCommand{UI: ui, Now: time.Now}
	cases := []struct {
		args     []string
		expected string
	}{
		{nil, ""},
		{[]string{"abc"}, "Invalid issue number"},
		{[]string{"-date=1234-56-78", "123"}, "Error, parsing time"},
		{[]string{"-date=1234", "123"}, "Error, parsing time"},
		{[]string{"-directory=invalid", "123"}, "no such file or directory"},
	}

	for _, c := range cases {
		code := cmd.Run(c.args)
		if code != 1 {
			t.Fatalf("Expected code = 1, got: %d (args = %v)", code, c.args)
		}
		if !strings.Contains(ui.ErrorWriter.String(), c.expected) {
			t.Fatalf(
				"Expected UI to contain '%s', got: %s (args = %v)",
				c.expected,
				ui.ErrorWriter.String(),
				c.args,
			)
		}

		ui.ErrorWriter.Reset()
	}
}

func TestNew(t *testing.T) {
	ui := new(cli.MockUi)
	dir := path.Join(os.TempDir(), "msw-test")
	os.Mkdir(dir, 0700)
	defer os.RemoveAll(dir)
	cases := []struct {
		now      func() time.Time
		date     string
		number   string
		expected string
	}{
		{time.Now, "2016-10-16", "123", "2016-10-16"},
		{fakeNow(2016, time.July, 17), "", "123", "2016-07-20"}, // Sunday
		{fakeNow(2016, time.July, 18), "", "123", "2016-07-20"}, // Monday
		{fakeNow(2016, time.July, 19), "", "123", "2016-07-20"}, // Tuesday
		{fakeNow(2016, time.July, 20), "", "123", "2016-07-20"}, // Wednesday
		{fakeNow(2016, time.July, 21), "", "123", "2016-07-27"}, // Thursday
	}

	for _, c := range cases {
		cmd := &command.NewCommand{UI: ui, Now: c.now}
		code := cmd.Run([]string{
			fmt.Sprintf("-directory=%s", dir),
			fmt.Sprintf("-date=%s", c.date),
			c.number,
		})
		if code != 0 {
			t.Fatalf("Command should return 0, got: %d", code)
		}

		data, err := ioutil.ReadFile(fmt.Sprintf("%s/%s.yml", dir, c.expected))
		if err != nil {
			t.Fatalf("A file should be generated: %s.yml (%s)", c.expected, err)
		}
		if !strings.Contains(string(data), fmt.Sprintf("number: %s", c.number)) {
			t.Fatalf("The generated file should contain '%s', got: %s", c.number, data)
		}

		// cleanup to ensure other tests with the same expected file pass or fail
		os.Remove(fmt.Sprintf("%s/%s.yml", dir, c.expected))
	}
}

func fakeNow(year int, month time.Month, day int) func() time.Time {
	return func() time.Time {
		return time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
	}
}
