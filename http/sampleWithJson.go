package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type lw struct {
	data []byte
}

func (l *lw) Write(p []byte) (n int, err error) {
	l.data = p
	return len(p), nil
}

type User struct {
	Userid    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Printf("Error %v", err)
		os.Exit(1)
	}
	logW := &lw{}

	io.Copy(logW, resp.Body)
	user := &User{}
	err = json.Unmarshal(logW.data, user)
	fmt.Println(user)
	re, err := json.Marshal(user)
	fmt.Println(string(re))
}
