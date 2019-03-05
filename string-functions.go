package main

import (
	"fmt"
	s "strings"
)

func main() {

	var p = fmt.Println

	var err error

	_, err = p("Contains", s.Contains("test", "es"))
	_, err = p("Count", s.Count("test", "s"))
	_, err = p("HasPrefix", s.HasPrefix("test", "te"))
	_, err = p("HasSuffix", s.HasSuffix("test", "st"))
	_, err = p("Index", s.Index("test", "s"))
	_, err = p("Join", s.Join([]string{"aa", "bb", "cc"}, "-"))
	_, err = p("Repeat", s.Repeat("aa", 3))
	_, err = p("Replace", s.Replace("test", "es", "ex", 1))
	_, err = p("Split", s.Split("aa-bb-cc-dd", "-"))
	_, err = p("ToLower", s.ToLower("Test"))
	_, err = p("ToUpper", s.ToUpper("teSt"))
	if err != nil {
		panic(err)
	}
}
