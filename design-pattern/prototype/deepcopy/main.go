package main

import "fmt"

type Address struct {
	StreetAdress, City, Country string
}

type Person struct {
	Name    string
	Address *Address
	Friends []string
}

func (a *Address) DeepCopy() *Address {
	return &Address{
		a.StreetAdress,
		a.City,
		a.Country,
	}
}

func (p *Person) DeepCopy() *Person {
	q := *p
	q.Address = p.Address.DeepCopy()
	copy(q.Friends, p.Friends)
	return &q
}
func main() {
	johny := Person{"Johny",
		&Address{"111 R Road", "London", "UK"}, []string{"Cat", "Mark"}}

	sarah := johny
	sarah.Name = "jinny"
	sarah.Address.StreetAdress = "4443 A Road"

	fmt.Println(johny, johny.Address) // Johny's Address also changed !
	fmt.Println(sarah, sarah.Address)

	jinny := johny
	jinny.Name = "Jinny"
	jinny.Address = &Address{
		johny.Address.StreetAdress,
		johny.Address.City,
		johny.Address.Country,
	}
	jinny.Address.StreetAdress = "4412 A Road"
	fmt.Println(johny, johny.Address)
	fmt.Println(jinny, jinny.Address)

	jane := johny.DeepCopy()
	jane.Name = "jane"
	jane.Address.StreetAdress = "31124 f road"
	jane.Friends = append(jane.Friends, "Buff")

	fmt.Println(jane, jane.Address)

}
