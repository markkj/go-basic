package main

import "fmt"

type Person struct {
	StreetAddress, Postcode, City string
	CompanyName, Position         string
	AnnualIncome                  int
}

type PersonBuilder struct {
	person *Person
}

func (b *PersonBuilder) Lives() *PersonAddressBuilder {
	return &PersonAddressBuilder{
		PersonBuilder: *b,
	}
}

func (b *PersonBuilder) Works() *PersonJobBuilder {
	return &PersonJobBuilder{
		PersonBuilder: *b,
	}
}

func NewPersonBuilder() *PersonBuilder {
	return &PersonBuilder{
		person: &Person{},
	}
}

type PersonAddressBuilder struct {
	PersonBuilder
}

type PersonJobBuilder struct {
	PersonBuilder
}

func (it *PersonAddressBuilder) At(StreetAddress string) *PersonAddressBuilder {
	it.person.StreetAddress = StreetAddress
	return it
}

func (it *PersonAddressBuilder) In(City string) *PersonAddressBuilder {
	it.person.City = City
	return it
}

func (it *PersonAddressBuilder) WithPostCode(Postcode string) *PersonAddressBuilder {
	it.person.Postcode = Postcode
	return it
}

func (jb *PersonJobBuilder) AsA(Position string) *PersonJobBuilder {
	jb.person.Position = Position
	return jb
}

func (jb *PersonJobBuilder) WorkAt(CompanyName string) *PersonJobBuilder {
	jb.person.CompanyName = CompanyName
	return jb
}

func (jb *PersonJobBuilder) Earnings(AnnualIncome int) *PersonJobBuilder {
	jb.person.AnnualIncome = AnnualIncome
	return jb
}

func (b *PersonBuilder) Build() *Person {
	return b.person
}

func main() {
	p := NewPersonBuilder()
	p.
		Lives().At("1234 B Road").In("London").WithPostCode("10230").
		Works().WorkAt("SomeCompany").AsA("Software Engineer").Earnings(100000)
	person := p.Build()
	fmt.Println(person)
}
