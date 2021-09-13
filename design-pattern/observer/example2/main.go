package main

import (
	"container/list"
	"fmt"
)

type Observable struct {
	subs *list.List
}

func (o *Observable) Subscribe(x Observer) {
	o.subs.PushBack(x)
}

func (o *Observable) UnSubscribe(x Observer) {
	for z := o.subs.Front(); z != nil; z = z.Next() {
		if z.Value.(Observer) == x {
			o.subs.Remove(z)
		}
	}
}

func (o *Observable) Fire(data interface{}) {
	for z := o.subs.Front(); z != nil; z = z.Next() {
		z.Value.(Observer).Notify(data)
	}
}

type Observer interface {
	Notify(data interface{})
}

type PropertyChange struct {
	Name  string
	Value interface{}
}

type Person struct {
	Observable
	age int
}

func NewPerson(age int) *Person {
	return &Person{
		Observable: Observable{new(list.List)},
		age:        age,
	}
}

func (p *Person) Age() int {
	return p.age
}

func (p *Person) SetAge(age int) {
	if age == p.age {
		return
	}

	oldCanVote := p.CanVote()

	p.age = age
	p.Fire(PropertyChange{"Age", p.age})

	if oldCanVote != p.CanVote() {
		p.Fire(PropertyChange{"CanVote", p.CanVote()})
	}
	fmt.Println("Setting Age to ", p.age)

}

func (p *Person) CanVote() bool {
	return p.age >= 18
}

type TrafficManaged struct {
	o Observable
}

func (t *TrafficManaged) Notify(data interface{}) {
	if pc, ok := data.(PropertyChange); ok {
		if pc.Value.(int) >= 18 {
			fmt.Println("Congrats, you can drive now !")
			t.o.UnSubscribe(t)
		}
	}
}

type ElectoralRoll struct {
	o Observable
}

func (e *ElectoralRoll) Notify(data interface{}) {
	if pc, ok := data.(PropertyChange); ok {
		if pc.Name == "CanVote" && pc.Value.(bool) {
			fmt.Println("Congrats, You can vote !")
		}
	}
}

func main() {
	p := NewPerson(17)
	t := &TrafficManaged{p.Observable}
	p.Subscribe(t)

	er := &ElectoralRoll{}
	p.Subscribe(er)
	p.UnSubscribe(er)
	for i := 16; i <= 19; i++ {
		p.SetAge(i)
	}

}
