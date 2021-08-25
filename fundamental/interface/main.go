package main

import "fmt"

type user struct {
	name string
}

type bot interface {
	getGreeting() string
	getName() string
}

type aBot struct{}

type bBot struct{}

func main() {
	a := aBot{}
	b := bBot{}

	printGreeting(a)
	printGreeting(b)
}

func printGreeting(bot bot) {
	fmt.Println(bot.getGreeting())
}

func (a aBot) getGreeting() string {
	return a.getName() + " Hello"
}
func (b bBot) getGreeting() string {
	return b.getName() + " Hello"
}

func (a aBot) getName() string {
	return "My name is A!"
}
func (b bBot) getName() string {
	return "My name is B!"
}
