package main

import (
	"fmt"
	"math/rand"
	"time"
)

// this tutorial in fucntion doSomethingA it maybe print or not print
func main() {
	c := make(chan string)
	go doSomethingA()
	go doSomethingB(c)
	go doSomethingC(c)
	fmt.Println(<-c)
	fmt.Println(<-c)

}

func doSomethingA() {
	time.Sleep(randomeTime() * time.Second)
	fmt.Println("Do Something in A")

}

func doSomethingB(message chan<- string) {
	time.Sleep(randomeTime() * time.Second)
	message <- "Do Something in B"

}

func doSomethingC(message chan<- string) {
	time.Sleep(randomeTime() * time.Second)
	message <- "Do Something in C"

}

func randomeTime() time.Duration {
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(3)
	return time.Duration(n)
}
