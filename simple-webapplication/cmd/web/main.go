package main

import (
	"fmt"
	"net/http"

	"github.com/markkj/simple-webapplication/pkg/handlers"
)

const portNumber = ":8080"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Printf("Starting Application running on port %s\n", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}
