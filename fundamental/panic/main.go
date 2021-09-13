package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("start")
	panicker()
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	fmt.Println("end")

}
func panicker() {
	fmt.Println("about to panic")

	defer func() {
		err := recover()
		if err != nil {
			log.Println("Error:", err)
		}
	}()

	panic("Shomething bad")
	fmt.Println("done panicker")
}
