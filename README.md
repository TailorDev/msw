msw
===

[![Circle CI](https://circleci.com/gh/TailorDev/msw.svg?style=svg)](https://circleci.com/gh/TailorDev/msw)

**msw** is a small tool that powers the
[ModernScienceWeekly](https://tinyletter.com/ModernScienceWeekly) newsletter,
written in Go. This is a full rewrite of one of our internal tools we did in
order to learn the Go language, which we also decided to open source.


## Usage

[Install the `msw` command](#installation), then run it:

    $ msw
    usage: ModernScienceWeekly [--version] [--help] <command> [<args>]

    Available commands are:
        generate    generate HTML for Tinyletter from a YAML file
        new         create a new empty YAML file to prepare a new issue.
        validate    check that an issue is valid


## Installation

    $ go get github.com/TailorDev/msw


## Development and Testing

If you wish to work on msw itself, you'll first need [Go](https://golang.org)
installed (version 1.6+ is _required_). Make sure you have Go properly
[installed](https://golang.org/doc/install), including setting up your
[GOPATH](https://golang.org/doc/code.html#GOPATH).

Next, clone this repository into `$GOPATH/src/github.com/TailorDev/msw`, and
run:

    $ go install

You can run the test suite with the following command:

    $ go test ./... [-cover]


## License

msw is released under the MIT License. See the bundled [LICENSE](LICENSE.md)
file for details.
