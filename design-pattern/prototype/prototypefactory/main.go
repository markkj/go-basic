package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

type Address struct {
	StreetAdress, City string
}

type Employee struct {
	Name   string
	Office Address
}

func (e *Employee) DeepCopy() *Employee {
	b := bytes.Buffer{}
	p := gob.NewEncoder(&b)
	p.Encode(&e)

	d := gob.NewDecoder(&b)
	result := &Employee{}
	d.Decode(result)
	return result

}

var mainOffice = Employee{
	"", Address{"1123 R Road", "London"},
}

var auxOffice = Employee{
	"", Address{"1234 Z Road", "New York"},
}

func NewEmployee(proto *Employee, name string) *Employee {
	result := proto.DeepCopy()
	result.Name = name
	return result
}

func NewMainOfficeEmployee(name string) *Employee {
	return NewEmployee(&mainOffice, name)
}

func NewAuxOfficeEmployee(name string) *Employee {
	return NewEmployee(&auxOffice, name)
}

func main() {
	johny := NewMainOfficeEmployee("John")
	jinny := NewAuxOfficeEmployee("Jinny")

	fmt.Println(johny)
	fmt.Println(jinny)
}
