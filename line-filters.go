package main

import (
	"bufio"
	"fmt"
	"go-example/util"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		ucl := strings.ToUpper(scanner.Text())

		fmt.Println(ucl)
	}

	if err := scanner.Err(); err != nil {
		_, err := fmt.Fprintln(os.Stderr, "error:", err)
		util.Check(err)
		os.Exit(1)
	}
}
