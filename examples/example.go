package main

import (
	"log"
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

func Example_AndThen_UnwrapOr() {
	// get home path, and then get log path, or use /tmp/log
	_ = AndThen11(getHomePath(), func(t1 string) Result1[string] {
		return Ok1(t1 + "/.log")
	}).UnwrapOr("/tmp/log")
}

func Example_Inspect_InspectErr() {
	// get home path, log it, or log error
	_ = getHomePath().Inspect(func(t1 string) {
		log.Println("home:", t1)
	}).InspectErr(func(err error) {
		log.Println("error:", err)
	}).UnwrapOr("/tmp/log")
}

func Example_Expect_ExpectErr() {
	// get home path, or panic with `must get home` msg
	_ = getHomePath().Expect("must get home")

	// get home path with err, or panic with `must get err` msg
	_ = getHomePath().ExpectErr("must get err")
}

func main() {
}
