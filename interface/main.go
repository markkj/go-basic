package main

import "fmt"

type Bot interface {
	getGreeting() string
	getName() string
}

type aBot struct{}

type bBot struct{}

func main() {
	a := aBot{}
	b := bBot{}

	printGreeting(a)
	fmt.Println(b.getName())
}

func printGreeting(bot Bot) {
	fmt.Println(bot.getGreeting())
}

func (a aBot) getGreetingNew() string {
	return a.getName() + " Hello From new"
}

func (a aBot) getGreeting() string {
	return "A Hello"
}

func (a aBot) getName() string {
	return "My name is A!"
}
func (b bBot) getName() string {
	return "My name is B!"
}
