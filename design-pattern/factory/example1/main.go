package main

import "fmt"

type Person struct {
	Firstname string
	Lastname  string
	Age       int
	Single    bool //Default
}

// Factory
func NewPerson(Firstname string, Lastname string, Age int) *Person {
	if Age < 18 {
		//do something
	}
	return &Person{Firstname, Lastname, Age, true}
}

func main() {
	p := NewPerson("Johny", "Deph", 40)
	fmt.Println(*p)
}
