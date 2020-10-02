# VLP (Vault Lock Picker)

[![go report card](https://goreportcard.com/badge/github.com/kenkyu392/vlp)](https://goreportcard.com/report/github.com/kenkyu392/vlp)
[![license](https://img.shields.io/github/license/kenkyu392/vlp.svg)](LICENSE)

VLP is a CLI tool for extracting passwords stored in credentials.

## Installation

```console
go get -u github.com/kenkyu392/vlp/cmd/vlp
```

## Usage

```console
$ vlp -help
VLP is a CLI tool for extracting passwords stored in credentials.
Usage: vlp <options>
  -help
        Prints the help.
  -service string
        Service name.
  -username string
        User name.
  -version
        Prints the version.
$ vlp -service=your-app -username=your-name
{"password":"secret","service":"your-app","username":"your-name"}
```

## License

[MIT](LICENSE)
