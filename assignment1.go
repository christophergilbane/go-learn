package main

import (
	"fmt"

	"example.com/user/assignment1/morestrings"
)

var (
	a string
	b string
	c string
)

func main() {
	a = "the "
	b = "answer is "
	c = "42 "
	var d string
	d = morestrings.ReturnConcatBob(a + b + c)
	fmt.Println(d)
}
