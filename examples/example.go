package main

import (
	"fmt"
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
	// setup workdir: $HOME/work or /tmp
	{
		homePath := getHomePath()
		workDir := MapOr(Inspect(homePath, func(t string) {
			fmt.Printf("home path is: %s\n", t)
		}), "/tmp", func(t string) string {
			return t + "/work"
		})
		_ = workDir
	}
}
