package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

func calSquareWithCtxAndDeadline(ctx context.Context, s int) {
	select {
	case <-ctx.Done():
		fmt.Printf("Job killed in: %v seconds \n", s)
		return
	case <-time.After(time.Duration(s) * (time.Second)):
		fmt.Printf("Job finish in: %v seconds \n", s)
	}

}
func main() {
	deadline := time.Now().Add(3 * time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel()
	for i := 0; i < 5; i++ {
		go calSquareWithCtxAndDeadline(ctx, i)
	}
  
	fmt.Println("Before: activate goroutine: ", runtime.NumGoroutine())
	time.Sleep(5 * time.Second)
	fmt.Println("After: activate goroutine: ", runtime.NumGoroutine())
}
