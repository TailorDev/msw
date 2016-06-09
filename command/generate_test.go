package command

import "testing"

func TestGenerateNoArgs(t *testing.T) {

	c := &GenerateCommand{Ui: mockUi()}

	code := c.Run(nil)
	if code != 1 {
		t.Fatalf("bad: %#v", code)
	}
}
