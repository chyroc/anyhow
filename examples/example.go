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
		workDir := getHomePath().UnwrapOr("/tmp")
		_ = workDir
	}
}
