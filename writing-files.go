package main

import (
	"bufio"
	"fmt"
	"go-example/util"
	"io/ioutil"
	"os"
)

func main() {
	d1 := []byte("hello\ngo\n")

	err := ioutil.WriteFile("/tmp/dat1", d1, 0644)
	util.Check(err)

	f, err := os.Create("/tmp/dat2")
	util.Check(err)

	defer f.Close()

	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := f.Write(d2)
	util.Check(err)
	fmt.Printf("wrote %d bytes\n", n2)

	n3, err := f.WriteString("writes\n")
	fmt.Printf("wrote %d bytes\n", n3)

	f.Sync()
	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered\n")
	fmt.Printf("wrote %d bytes\n", n4)
	w.Flush()
}