package main

import "fmt"

func main() {
	fmt.Println(normalFunction("hello "))
	f := func(a string) string {
		return a + "inside"
	}("hello ")
	fmt.Println(f)

}

func normalFunction(a string) string {
	return a + "normal"
}
