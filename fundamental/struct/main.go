package main

import "fmt"

type person struct {
	firstname string
	lastname  string
}

type author struct {
	startDate string
	person
}

type blogPost struct {
	title string
	author
}

func main() {
	peter := &person{
		firstname: "Peter",
		lastname:  "Jologo",
	}
	author1 := &author{
		person: *peter,
	}
	firstPost := &blogPost{
		title:  "firstPost",
		author: *author1,
	}
	firstPost.printDetails()
}

func (b blogPost) printDetails() {
	fmt.Println("title :", b.title)
	b.author.printDetails()
}

func (p *person) printDetails() {
	fmt.Println("fistname :", p.firstname)
	fmt.Println("lastname :", p.lastname)
}
