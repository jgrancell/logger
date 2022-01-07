logger
===

[![GoDoc](https://godoc.org/github.com/jgrancell/logger?status.svg)](https://pkg.go.dev/github.com/jgrancell/logger)
[![Go Report Card](https://goreportcard.com/badge/jgrancell/logger)](https://goreportcard.com/report/jgrancell/logger)
[![codecov](https://codecov.io/gh/jgrancell/logger/branch/master/graph/badge.svg)](https://codecov.io/gh/jgrancell/logger)

Logger is a small, straightforward library to enhance your Go applications with powerful logging features.

## Overview
Logger is a library providing a simple interface to add feature rich and customizable logging support to your application.

Logger is composed of modular Handlers, each of which implement a different mechanism by which your application can output logging information.

Currently supported handlers include:
- Standard Out / Standard Error
- File Logging

Logger also includes a number of modular Formatters, which transform your log messages into standards-compliant formats before they are logged.

Currently supported formatters include:
- [Logger Structured Log Format](./docs/formats.md#logger-structured-log-format)
- JSON
- Graylog Extended Format
- Kubernetes Structured Log Format
-

## Usage Documentation



## Installation

Using this package requires a working Go environment. [See the install instructions for Go](http://golang.org/doc/install.html).

```go
...
import (
  "github.com/jgrancell/logger"
)
...
```

### GOPATH

Make sure your `PATH` includes the `$GOPATH/bin` directory so your commands can
be easily used:
```
export PATH=$PATH:$GOPATH/bin
```

### Supported platforms

cli is tested against multiple versions of Go on Linux, and against the latest
released version of Go on OS X and Windows. This project uses Github Actions for
builds. To see our currently supported go versions and platforms, look at the [./.github/workflows/cli.yml](https://github.com/urfave/cli/blob/master/.github/workflows/cli.yml).