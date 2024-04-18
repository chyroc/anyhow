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

## Features

- result struct
  - Result1[T]
  - Result2[T1, T2]
  - Result3[T1, T2, T3]
  - Result4[T1, T2, T3, T4]
  - Result5[T1, T2, T3, T4, T5]
  - Result6[T1, T2, T3, T4, T5, T6]
- functions:
  - And(self R[T], res R[U]) R[U]: Returns res if the result is Ok, otherwise returns the Err value of self
  - AndThen(self R[T], op (T) -> R[U]) R[U]: Calls op if the result is Ok, otherwise returns the Err value of self
  - Map(self R[T], op (T) -> U) R[U]: Maps a Result<T> to Result<U> by applying a function to a contained Ok value, leaving an Err value untouched
  - MapOr(self R[T], default U, op func(T) U) U: Returns the provided default (if Err), or applies a function to the contained value (if Ok)
- methods:
  - (R[T]) IntoErr() error: Returns the contained Err value, but never panics
  - (R[T]) IntoOk() T1: Returns the contained Ok value, but never panics
  - (R[T]) Expect(string) T: Returns the contained Ok value, otherwise panics with a panic message including the passed message, and the content of the Err
  - (R[T]) ExpectErr(string) error: Returns the contained Err value, otherwise panics with a panic message including the passed message, and the content of the Ok
  - (R[T]) Inspect(f (T) -> void) R[T]: Calls a function with a reference to the contained value if Ok
  - (R[T]) InspectErr(f (error) -> void) R[T]: Calls a function with a reference to the contained value if Err
  - (R[T]) MapErr(op func(error) error) R[T]: Maps a R[T] to R[T] by applying a function to a contained Err value, leaving an Ok value untouched
  - (R[T]) IsErr() bool: Returns true if the result is Err
  - (R[T]) IsOk() bool: Returns true if the result is Ok
  - (R[T]) Or(res R[T]) R[T]: Returns res if the result is Err, otherwise returns the Ok value of self
  - (R[T]) OrElse(op (error) -> R[T]) R[T]: Calls op if the result is Err, otherwise returns the Ok value of self
  - (R[T]) Unwrap() T: Returns the contained Ok value, otherwise panics with Err
  - (R[T]) UnwrapOr(default T) T: Returns the contained Ok value or a provided default
  - (R[T]) UnwrapOrDefault() T: Returns the contained Ok value or a default value
  - (R[T]) UnwrapOrElse(op (error) -> T) T: Returns the contained Ok value or computes it from a closure

## Usage

```go
package main

import (
	"os/user"

	. "github.com/chyroc/anyhow"
)

func getHomePath() Result1[string] {
	user, err := user.Current()
	if err != nil {
		return Err1[string](err)
	}
	return Ok1(user.HomeDir)
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

