package main

import (
	"fmt"
	"time"
)

func main() {

	timer1 := time.NewTimer(time.Second * 2)

	<-timer1.C
	fmt.Println("Timer 1 expired")

}
