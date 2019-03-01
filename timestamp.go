package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()

	secs := now.Unix()
	nsec := now.UnixNano()

	fmt.Println(now)

	millis := nsec / 1000000
	fmt.Println(secs)
	fmt.Println(millis)
	fmt.Println(nsec)

	fmt.Println(time.Unix(secs, 0))
	fmt.Println(time.Unix(0, nsec))
}
