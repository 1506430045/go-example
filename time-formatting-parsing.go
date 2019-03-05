package main

import (
	"fmt"
	"time"
)

var p = fmt.Println

func main() {
	t := time.Now()
	var err error

	_, err = p(t.Format(time.RFC3339))

	t1, err := time.Parse(
		time.RFC3339,
		"2012-11-01T22:08:41+00:00",
	)

	_, err = p(t1)
	_, err = p(t.Format("3:04PM"))
	_, err = p(t.Format("Mon Jan _2 15:04:05 2006"))
	_, err = p(t.Format("2006-01-02T15:04:05.999999-07:00"))
	from := "3 04 PM"
	t2, err := time.Parse(from, "8 41 PM")
	_, err = p(t2)

	fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())

	ansic := "Mon Jan _2 15:04:05 2006"
	_, err = time.Parse(ansic, "8:41PM")

	if err != nil {
		panic(err)
	}
}
