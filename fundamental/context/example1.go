package main

import (
	"fmt"
	"runtime"
	"time"
)

func someSlowWorkFnc(c chan string) {
	for i := 0; i < 3; i++ {
		time.Sleep(2 * time.Second)

	}
	c <- "Finish"
}

func calSquare(c chan int) {
	i := 1
	for {
		c <- i * i
		i++
	}
}

func main() {
	c := make(chan int, 1)
	go calSquare(c)
	fmt.Println("Before: activate goroutine: ", runtime.NumGoroutine())
	for i := 0; i < 5; i++ {
		fmt.Println(<-c)
	}

	time.Sleep(3 * time.Second)
	fmt.Println("After: activate goroutine: ", runtime.NumGoroutine())
}
