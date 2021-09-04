package main

import "fmt"

var (
	a int = 10
	b     = "ss"
)

const (
	c = 10
	d = "hello"
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	a += 10
	b += "ff"

	fmt.Println(a)
	fmt.Println(b)

}
