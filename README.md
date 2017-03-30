msw
===

[![Circle CI](https://circleci.com/gh/TailorDev/msw.svg?style=svg)](https://circleci.com/gh/TailorDev/msw)

**msw** is a small tool that powers the
[ModernScienceWeekly](https://tinyletter.com/ModernScienceWeekly) newsletter,
written in Go. This is a full rewrite of one of our internal tools we did in
order to learn the Go language, which we also decided to open source. You can
find our blog post about this very first experience:

* https://tailordev.fr/blog/2016/06/23/a-tour-of-go/


## Usage

1. [Install the `msw` command](#installation)
2. (optionally) create a configuration file `~/.msw/msw.toml`:

    ``` toml
    # ~/.msw/msw.toml
    [buffer]
    AccessToken = "BUFFER_ACCESS_TOKEN"
    ProfileIDs  = ["BUFFER_PROFILE_ID"]
    ```

3. Run `msw`:

    ```
    $ msw
    usage: msw [--version] [--help] <command> [<args>]

    Available commands are:
        buffer
        generate    generate HTML for Tinyletter from a YAML file
        new         create a new empty YAML file to prepare a new issue
        validate    check that an issue is valid
    ```


## Installation

    $ go get github.com/TailorDev/msw

Alternatively, you can download pre-compiled packages: https://github.com/TailorDev/msw/releases.

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

You can build packages with [gox](https://github.com/mitchellh/gox):

    $ go get github.com/mitchellh/gox
    ...
    $ gox -osarch="darwin/amd64 linux/amd64" -output="pkg/msw_{{.OS}}_{{.Arch}}"
    Number of parallel builds: 1

    -->    darwin/amd64: github.com/TailorDev/msw
    -->     linux/amd64: github.com/TailorDev/msw


## License

msw is released under the MIT License. See the bundled [LICENSE](LICENSE.md)
file for details.
