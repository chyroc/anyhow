# anyhow

go impl for https://docs.rs/anyhow/latest/anyhow/

[![codecov](https://codecov.io/gh/chyroc/anyhow/branch/master/graph/badge.svg)](https://codecov.io/gh/chyroc/anyhow)
[![go report card](https://goreportcard.com/badge/github.com/chyroc/anyhow "go report card")](https://goreportcard.com/report/github.com/chyroc/anyhow)
[![test status](https://github.com/chyroc/anyhow/actions/workflows/test.yml/badge.svg)](https://github.com/chyroc/anyhow/actions)
[![Apache-2.0 license](https://img.shields.io/badge/License-Apache%202.0-brightgreen.svg)](https://opensource.org/licenses/Apache-2.0)
[![Go.Dev reference](https://img.shields.io/badge/go.dev-reference-blue?logo=go&logoColor=white)](https://pkg.go.dev/github.com/chyroc/anyhow)
[![Go project version](https://badge.fury.io/go/github.com%2Fchyroc%2Fanyhow.svg)](https://badge.fury.io/go/github.com%2Fchyroc%2Fanyhow)

## Installation

```shell
go get github.com/chyroc/anyhow
```

## Usage

```go
package main

import (
	"os/user"

	. "github.com/chyroc/anyhow"
)

type number interface {
	~int | ~int32
}

func addOne[T number](x T) T {
	return x + 1
}

func getHomePath() Result[string] {
	user, err := user.Current()
	if err != nil {
		return AErr[string](err)
	}
	return AOk(user.HomeDir)
}

func main() {
	// setup workdir: $HOME or /tmp
	{
		workDir := getHomePath().
			UnwrapOr("/tmp")
		_ = workDir
	}
}
```

