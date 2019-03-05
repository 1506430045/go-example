package main

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"
)

func main() {
	s := "sha1 this string"

	fmt.Println(Sha1(s))

	s = "md5 this string"

	fmt.Println(Md5(s))

}

func Sha1(str string) string {
	h := sha1.New()
	h.Write([]byte(str))
	bs := h.Sum(nil)

	return fmt.Sprintf("%x", bs)
}

func Md5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	bs := h.Sum(nil)

	return fmt.Sprintf("%x", bs)
}
