package command_test

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"github.com/TailorDev/msw/command"
	"github.com/mitchellh/cli"
)

func TestNewInvalidArgs(t *testing.T) {
	ui := new(cli.MockUi)
	cmd := &command.NewCommand{UI: ui}
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
	c := &command.NewCommand{UI: ui}
	dir := os.TempDir()

	code := c.Run([]string{
		fmt.Sprintf("-directory=%s", dir),
		"-date=2016-10-16",
		"123",
	})
	if code != 0 {
		t.Fatalf("Command should return 0, got: %d", code)
	}

	data, err := ioutil.ReadFile(fmt.Sprintf("%s/2016-10-16.yml", dir))
	if err != nil {
		t.Fatalf("A file should be generated: %s", err)
	}
	if !strings.Contains(string(data), "number: 123") {
		t.Fatalf("The generated file should contain '123', got: %s", data)
	}
}
