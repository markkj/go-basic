package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.orgfaf",
		"http://amazon.com",
	}
	status := make([]string, len(links))
	c := make(chan []string)

	for i, link := range links {
		go checkLink(link, i, c)
	}

	for i := 0; i < len(status); i++ {
		data := <-c
		index, _ := strconv.Atoi(data[1])
		status[index] = links[index] + data[0]
	}
	fmt.Println(status)
}

func checkLink(link string, index int, c chan []string) {
	da, err := http.Get(link)
	if err != nil {
		c <- []string{"migth be down", strconv.Itoa(index)}
		return
	}
	// fmt.Println(da.Status, link)
	c <- []string{da.Status, strconv.Itoa(index)}
	return
}
