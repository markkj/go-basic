package main

import "fmt"

type Employee struct {
	Name, Position string
	Income         int
}

const (
	JuniorDev = iota
	SeniorDev
	Techlead
)

func NewEmployee(role int) *Employee {
	switch role {
	case JuniorDev:
		return &Employee{"", "junior dev", 5000}
	case SeniorDev:
		return &Employee{"", "senior dev", 10000}
	case Techlead:
		return &Employee{"", "Technical Lead", 20000}
	default:
		panic("Not found this role")
	}
}

func main() {
	juniorDev := NewEmployee(JuniorDev)
	seniorDev := NewEmployee(SeniorDev)
	techLead := NewEmployee(Techlead)

	juniorDev.Name = "M"
	seniorDev.Name = "C"
	techLead.Name = "Z"
	fmt.Println(juniorDev)
	fmt.Println(seniorDev)
	fmt.Println(techLead)

}
