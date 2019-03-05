package main

import (
	"fmt"
	"go-example/util"
	"os"
	"strings"
)

func main() {
	err := os.Setenv("FOO", "1")
	util.Check(err)

	fmt.Println("FOO:", os.Getenv("FOO"))
	fmt.Println("BAR:", os.Getenv("BAR"))

	fmt.Println()

	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		fmt.Println(pair[0])
	}
}
