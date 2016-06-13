package command_test

import (
	"bytes"

	"github.com/mitchellh/cli"
)

func mockUi() cli.Ui {
	var out, err bytes.Buffer

	return &cli.BasicUi{
		Writer:      &out,
		ErrorWriter: &err,
	}
}
