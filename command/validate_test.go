package command

import "testing"

func TestValidateNoArgs(t *testing.T) {

	c := &ValidateCommand{Ui: mockUi()}

	code := c.Run(nil)
	if code != 1 {
		t.Fatalf("bad: %#v", code)
	}
}
