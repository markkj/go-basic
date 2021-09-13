package main

import (
	"fmt"
	"io"
	"os"
)

type logWritter struct{}

func main() {
	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Get some err %v", err)
		os.Exit(1)
	}
	lw := logWritter{}
	io.Copy(lw, file)
}

func (logWritter) Write(bs []byte) (n int, err error) {
	fmt.Println(string(bs))
	return len(bs), nil
}
