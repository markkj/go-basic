package main

import "fmt"

type Person interface {
	sayHello()
}

type person struct {
	name string
	age  int
}

type oldPerson struct {
	name string
	age  int
}

func (p *person) sayHello() {
	fmt.Printf("Hello, my name is %s, I'm %d years old \n", p.name, p.age)
}

func (p *oldPerson) sayHello() {
	fmt.Println("Hello, I'm old ")
}

func NewPerson(name string, age int) Person {
	if age > 200 {
		return &oldPerson{name, age}
	}
	return &person{name, age}
}

func main() {
	p := NewPerson("Johny", 42)
	p.sayHello()

	oldP := NewPerson("Lucas", 400)
	oldP.sayHello()
}
