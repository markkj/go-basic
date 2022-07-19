package main

import (
	"log"
	"net/http"
)

type routes []func(w http.ResponseWriter, r *http.Request)

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}

func main() {
	rs := routes{
		HelloWorld,
	}
	for _, r := range rs {
		http.HandleFunc("/test", r)
	}

	err := http.ListenAndServe(":3333", nil)
	if err != nil {
		log.Fatalf("error: %+v", err)
	}
}
