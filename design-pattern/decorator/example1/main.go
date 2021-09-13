package main

import "fmt"

type Aged interface {
	GetAge() int
	SetAge(age int)
}

type Bird struct {
	Age int
}

func (b *Bird) GetAge() int {
	return b.Age
}
func (b *Bird) SetAge(age int) {
	b.Age = age
}

func (b *Bird) Fly() {
	if b.Age >= 10 {
		fmt.Println("Flying")
	}
}

type Lizard struct {
	Age int
}

func (l *Lizard) GetAge() int {
	return l.Age
}
func (l *Lizard) SetAge(age int) {
	l.Age = age
}

func (l *Lizard) Crawl() {
	if l.Age < 10 {
		fmt.Println("Crawling")
	}
}

type Dragon struct {
	Bird
	Lizard
}

func (d *Dragon) GetAge() int {
	return d.Bird.Age
}
func (d *Dragon) SetAge(age int) {
	d.Lizard.Age = age
	d.Bird.Age = age
}

func main() {
	d := Dragon{}
	d.SetAge(10)
	d.Fly()
	d.Crawl()

}
