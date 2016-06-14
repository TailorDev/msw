package command_test

import (
	"strings"
	"testing"

	"github.com/TailorDev/msw/command"
	"github.com/mitchellh/cli"
)

func TestGenerateNoArgs(t *testing.T) {
	c := &command.GenerateCommand{UI: new(cli.MockUi)}

	code := c.Run(nil)
	if code != 1 {
		t.Fatalf("Expected code = 1, got: %d", code)
	}
}

func TestGenerate(t *testing.T) {
	ui := new(cli.MockUi)
	c := &command.GenerateCommand{UI: ui}

	code := c.Run([]string{"../test-fixtures/2016-10-13.yml"})
	if code != 0 {
		t.Fatalf("TestGenerate should work correctly")
	}
	if !strings.Contains(ui.OutputWriter.String(), "Issue #123 &mdash; 10/13/2016") {
		t.Fatalf("got: %s", ui.OutputWriter)
	}
}

func BenchmarkGenerate(b *testing.B) {
	c := &command.GenerateCommand{UI: new(cli.MockUi)}

	for i := 0; i < b.N; i++ {
		code := c.Run([]string{"../test-fixtures/2016-10-13.yml"})
		if code != 0 {
			b.Fatalf("TestGenerate should work correctly")
		}
	}
}
