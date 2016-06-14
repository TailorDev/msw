package command_test

import (
	"testing"

	"github.com/TailorDev/msw/command"
	"github.com/mitchellh/cli"
)

func TestValidateNoArgs(t *testing.T) {

	c := &command.ValidateCommand{UI: new(cli.MockUi)}

	code := c.Run(nil)
	if code != 1 {
		t.Fatalf("Expected code = 1, got: %d", code)
	}
}
