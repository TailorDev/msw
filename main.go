package main

import (
	"fmt"
	"os"
	"time"

	"github.com/TailorDev/msw/command"
	"github.com/TailorDev/msw/config"
	"github.com/TailorDev/msw/version"
	"github.com/mitchellh/cli"
)

func main() {
	ui := &cli.BasicUi{
		Reader:      os.Stdin,
		Writer:      os.Stdout,
		ErrorWriter: os.Stderr,
	}

	c := cli.NewCLI("msw", version.FormattedVersion())
	c.Args = os.Args[1:]

	conf, err := config.LoadDefaultConfig()
	if err != nil {
		fmt.Fprintf(os.Stderr, "WARNING: %s\n", err)
	}

	c.Commands = map[string]cli.CommandFactory{
		"generate": func() (cli.Command, error) {
			return &command.GenerateCommand{UI: ui}, nil
		},
		"new": func() (cli.Command, error) {
			return &command.NewCommand{UI: ui, Now: time.Now}, nil
		},
		"validate": func() (cli.Command, error) {
			return &command.ValidateCommand{UI: ui}, nil
		},
		"buffer push": func() (cli.Command, error) {
			return &command.BufferPushCommand{UI: ui, Conf: *conf}, nil
		},
		"buffer info": func() (cli.Command, error) {
			return &command.BufferInfoCommand{UI: ui, Conf: *conf}, nil
		},
	}

	exitStatus, err := c.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error executing CLI: %s\n", err)
	}

	os.Exit(exitStatus)
}
