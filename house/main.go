package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

type stringArray []string

func main() {
	p1 := person{
		firstName: "Khajohnyos",
		lastName:  "Boonkate",
		contactInfo: contactInfo{
			email:   "mark@gmail.com",
			zipCode: 1204,
		},
	}
	p1.print()
	fmt.Println(&p1)
	p1.updateName("Mark")
	p1.print()

}

func (p *person) updateName(firstName string) {
	(*p).firstName = firstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
	fmt.Println("")
}
