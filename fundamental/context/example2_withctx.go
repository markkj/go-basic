package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

func calSquareWithCtx(ctx context.Context, c chan int) {
	i := 1
	for {
		select {
		case <-ctx.Done():
			return
		default:
			c <- i * i
			i++
		}
	}
}

func main() {
	c := make(chan int, 1)

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	go calSquareWithCtx(ctx, c)
	fmt.Println("Before: activate goroutine: ", runtime.NumGoroutine())
	for i := 0; i < 5; i++ {
		fmt.Println(<-c)
	}
  cancel()
	time.Sleep(3 * time.Second)
	fmt.Println("After: activate goroutine: ", runtime.NumGoroutine())
}
