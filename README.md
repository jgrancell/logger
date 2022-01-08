logger
===

[![GoDoc](https://godoc.org/github.com/jgrancell/logger?status.svg)](https://pkg.go.dev/github.com/jgrancell/logger)
[![codebeat badge](https://codebeat.co/badges/19c16d5c-407f-478a-9b4e-3ffdb070607c)](https://codebeat.co/projects/github-com-jgrancell-logger-main)
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