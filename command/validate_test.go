package command_test

import (
	"testing"

	"github.com/TailorDev/msw/command"
)

func TestValidateNoArgs(t *testing.T) {

	c := &command.ValidateCommand{Ui: mockUi()}

	code := c.Run(nil)
	if code != 1 {
		t.Fatalf("bad: %#v", code)
	}
}
