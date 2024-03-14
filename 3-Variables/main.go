package main

import (
	"fmt"
)

func main() {
	a := "Hello world"
	fmt.Println(a)

	var b string
	b = "Hello world 2"
	fmt.Println(b)

	c, d, e := 1, "Hello", 3
	fmt.Println(c, d, e)

	var f, g, h int = 4, 5, 6
	fmt.Println(f, g, h)
}
