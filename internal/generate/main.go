//go:build ignore

package main

import (
	"os"
)

func main() {
	{
		f, err := os.Create("struct.go")
		assert(err)
		f.Write([]byte(generateStruct(6)))
		f.Close()
	}
	{
		f, err := os.Create("constructor.go")
		assert(err)
		f.Write([]byte(generateConstructor(6)))
		f.Close()
	}
	{
		f, err := os.Create("method.go")
		assert(err)
		f.Write([]byte(generateMethod(6)))
		f.Close()
	}
	{
		f, err := os.Create("func.go")
		assert(err)
		f.Write([]byte(generateFunc(6)))
		f.Close()
	}
}

func assert(err error) {
	if err != nil {
		panic(err)
	}
}
