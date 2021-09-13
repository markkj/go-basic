package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

type Address struct {
	StreetAdress, City, Country string
}

type Person struct {
	Name    string
	Address *Address
	Friends []string
}

func (p *Person) DeepCopy() *Person {
	b := bytes.Buffer{}
	e := gob.NewEncoder(&b)
	_ = e.Encode(p)

	d := gob.NewDecoder(&b)
	result := &Person{}
	_ = d.Decode(result)

	return result
}
func main() {
	johny := Person{"Johny",
		&Address{"111 R Road", "London", "UK"}, []string{"Cat", "Mark"}}
	jinny := johny.DeepCopy()
	jinny.Address.StreetAdress = "123 A Road"

	fmt.Println(johny, johny.Address)
	fmt.Println(jinny, jinny.Address)

}
