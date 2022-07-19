package main

import "fmt"

func squreVal(v *int) {
	*v *= *v
	fmt.Println("value", v, "address", &v)
}

func squreWithVal(v int) {
	v *= v
	fmt.Println("value", v, "address", &v)
}

func main() {
	x := 10
	fmt.Println(&x)
	squreWithVal(x)
	squreVal(&x)
	// fmt.Println(x, &x)
	// squreVal(&x)
}
