package command_test

import (
	"strings"
	"testing"

	"github.com/TailorDev/msw/command"
	"github.com/TailorDev/msw/config"
	"github.com/mitchellh/cli"
)

func TestBufferPushNoArgs(t *testing.T) {
	config := config.DefaultConfig
	c := &command.BufferPushCommand{
		UI:   new(cli.MockUi),
		Conf: config,
	}

	code := c.Run(nil)
	if code != 1 {
		t.Fatalf("Expected code = 1, got: %d", code)
	}
}

func TestBufferPush(t *testing.T) {
	ui := new(cli.MockUi)
	config := config.Config{
		Buffer: config.BufferOptions{
			AccessToken: "ACCESS_TOKEN",
			ProfileIDs:  []string{"PROFILE_ID"},
			BufferSize:  10,
		},
	}
	c := &command.BufferPushCommand{
		UI:   ui,
		Conf: config,
	}

	code := c.Run([]string{"../test-fixtures/2016-10-13.yml"})
	if code != 0 {
		t.Fatalf("Command should return 0, got: %d", code)
	}
	if !strings.Contains(ui.OutputWriter.String(), "[?] Link #1: http://example.org.") {
		t.Fatalf("got: %s", ui.OutputWriter)
	}
}
