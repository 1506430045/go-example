package main

import (
	"fmt"
	"strconv"
)

func main() {
	f, _ := strconv.ParseFloat("1.23", 64)
	fmt.Println(f)

	i, _ := strconv.ParseInt("323", 0, 64)
	fmt.Println(i)

	d, _ := strconv.ParseInt("0x1c8", 0, 64)
	fmt.Println(d)

	u, _ := strconv.ParseUint("333", 0, 64)
	fmt.Println(u)

	k, _ := strconv.Atoi("121")
	fmt.Println(k)

	str := strconv.Itoa(222)
	fmt.Println(str)
}
