package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWritter struct {
}

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Printf("Error %v", err)
		os.Exit(1)
	}
	// bs := make([]byte, 99999)
	// resp.Body.Read(bs)
	// fmt.Println(string(bs))
	// resp.Body.Close()

	lw := logWritter{}
	io.Copy(lw, resp.Body)

}

func (logWritter) Write(bs []byte) (int, error) {
	return len(bs), nil
}
