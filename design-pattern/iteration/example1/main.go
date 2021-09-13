package main

import "fmt"

type Person struct {
	FirstNanme, MiddleName, LastName string
}

func (p *Person) Names() [3]string {
	return [3]string{p.FirstNanme, p.MiddleName, p.LastName}
}

func (p *Person) NamesGenerator() <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		out <- p.FirstNanme
		if len(p.MiddleName) > 0 {
			out <- p.MiddleName
		}
		out <- p.LastName
	}()
	return out
}

func main() {
	p := Person{"Alexander", "Johny", "Jinny"}
	for name := range p.NamesGenerator() {
		fmt.Println(name)
	}
}
