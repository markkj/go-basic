package main

import "fmt"

type Employee struct {
	Name, Position string
	Income         int
}

func NewEmployeeFactory(position string, income int) func(name string) *Employee {
	return func(name string) *Employee {
		return &Employee{name, position, income}
	}
}

type EmployeeFactory struct {
	Position string
	Income   int
}

func (f *EmployeeFactory) Create(name string) *Employee {
	return &Employee{name, f.Position, f.Income}
}

func NewEmployeeFactory2(position string, income int) *EmployeeFactory {
	return &EmployeeFactory{position, income}
}

func main() {
	devFactory := NewEmployeeFactory("Dev", 100000)
	managerFactory := NewEmployeeFactory("Manager", 200000)

	johnDev := devFactory("John")
	jinnyManager := managerFactory("Jinny")

	fmt.Println(johnDev)
	fmt.Println(jinnyManager)

	//Another Implement
	CEOFactory := NewEmployeeFactory2("CEO", 1000000)
	SaraCEO := CEOFactory.Create("Sara")
	CEOFactory.Income = 1100000 //Can Change Income but not changes sara's income

	fmt.Println(SaraCEO)

}
