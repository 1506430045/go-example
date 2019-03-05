package main

import (
	"fmt"
	"os"
)

func main() {
	argsWithProgram := os.Args
	argsWithoutProgram := argsWithProgram[1:]

	arg := os.Args[3]

	fmt.Println(argsWithProgram)
	fmt.Println(argsWithoutProgram)
	fmt.Println(arg)
}
