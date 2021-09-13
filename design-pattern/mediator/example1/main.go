package main

import "fmt"

type Person struct {
	Name    string
	Room    *Chatroom
	chatLog []string
}

func NewPerson(name string) *Person {
	return &Person{Name: name}
}

func (p *Person) Receive(sender, message string) {
	s := fmt.Sprintf("%s: %s", sender, message)
	fmt.Printf("[%s's chat session]: %s \n", p.Name, s)
	p.chatLog = append(p.chatLog, s)
}

func (p *Person) Say(message string) {
	p.Room.Broadcast(p.Name, message)
}

func (p *Person) PrivateMessage(who, message string) {
	p.Room.Message(p.Name, who, message)
}

type Chatroom struct {
	people []*Person
}

func (c *Chatroom) Broadcast(src, message string) {
	for _, p := range c.people {
		if p.Name != src {
			p.Receive(src, message)
		}
	}
}

func (c *Chatroom) Message(src, dst, message string) {
	for _, p := range c.people {
		if p.Name != dst {
			p.Receive(src, message)
		}
	}
}

func (c *Chatroom) Join(p *Person) {
	msg := p.Name + " joins the room"
	c.Broadcast("Room", msg)
	p.Room = c
	c.people = append(c.people, p)
}

func main() {
	room := Chatroom{}

	mark := NewPerson("Mark")
	john := NewPerson("John")

	room.Join(mark)
	room.Join(john)

	john.Say("Hi Mark")
	mark.Say("Hello John")

	nicha := NewPerson("Nicha")
	room.Join(nicha)

	nicha.Say("Hello My name is Nicha")

}
