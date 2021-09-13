package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// this tutorial in fucntion doSomethingA it will print
var wg = sync.WaitGroup{}

func main() {
	c := make(chan string)
	wg.Add(1)
	go doSomethingA()
	go doSomethingB(c)
	go doSomethingC(c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	wg.Wait()

}

func doSomethingA() {
	time.Sleep(randomeTime() * time.Second)
	fmt.Println("Do Something in A")
	wg.Done()
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
