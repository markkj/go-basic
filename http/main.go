package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/markkj/go-basic/http/helpers"
)

type logWritter struct {
	write string
	done  bool
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

	lw := &logWritter{}
	// fmt.Println(lw)
	io.Copy(lw, resp.Body)
	// fmt.Println(lw)

	h := helpers.WEBURL{}
	h.WebURL = "http://google.com"
	h.WebStatus = resp.Status
	fmt.Println(h)

	newJson, err := json.Marshal(h)
	fmt.Println(string(newJson))
}

func (lw *logWritter) Write(bs []byte) (int, error) {
	lw.write = string(bs)
	return len(bs), nil
}
