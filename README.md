# Logur adapter for [hclog](https://github.com/hashicorp/go-hclog)

![GitHub Workflow Status](https://img.shields.io/github/workflow/status/logur/adapter-hclog/CI?style=flat-square)
[![Codecov](https://img.shields.io/codecov/c/github/logur/adapter-hclog?style=flat-square)](https://codecov.io/gh/logur/adapter-hclog)
[![Go Report Card](https://goreportcard.com/badge/logur.dev/adapter/hclog?style=flat-square)](https://goreportcard.com/report/logur.dev/adapter/hclog)
[![GolangCI](https://golangci.com/badges/github.com/logur/adapter-hclog.svg)](https://golangci.com/r/github.com/logur/adapter-hclog)
[![Go Version](https://img.shields.io/badge/go%20version-%3E=1.11-61CFDD.svg?style=flat-square)](https://github.com/logur/adapter-hclog)
[![GoDoc](http://img.shields.io/badge/godoc-reference-5272B4.svg?style=flat-square)](https://godoc.org/logur.dev/adapter/hclog)


## Installation

```bash
go get logur.dev/adapter/hclog
```


## Usage

```go
package main

import (
	"github.com/hashicorp/go-hclog"
	hclogadapter "logur.dev/adapter/hclog"
)

func main() {
	logger := hclogadapter.New(hclog.Default())
}
```


## Development

When all coding and testing is done, please run the test suite:

```bash
$ make check
```


## License

The MIT License (MIT). Please see [License File](LICENSE) for more information.
