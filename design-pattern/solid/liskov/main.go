package main

import (
	"fmt"
)

// Liskov Substitution Principle
// Sub Class should be perform like Parent Class

type Sized interface {
	GetWidth() int
	SetWidth(width int)
	GetHeight() int
	SetHeight(height int)
}

type Rectangle struct {
	width, height int
}

func (r *Rectangle) GetWidth() int {
	return r.width
}

func (r *Rectangle) SetWidth(width int) {
	r.width = width
}

func (r *Rectangle) GetHeight() int {
	return r.height
}

func (r *Rectangle) SetHeight(height int) {
	r.height = height
}

type Square struct {
	Rectangle
}

func NewSquare(size int) *Square {
	sq := Square{}
	sq.width = size
	sq.height = size
	return &sq
}

func (s *Square) GetWidth() int {
	return s.width
}

func (s *Square) SetWidth(width int) {
	s.width = width
	s.height = width
}

func (s *Square) GetHeight() int {
	return s.height
}

func (s *Square) SetHeight(height int) {
	s.width = height
	s.height = height
}

type Square2 struct {
	size int
}

func (s *Square2) Rectangle() Rectangle {
	return Rectangle{s.size, s.size}
}

func UseIt(sized Sized) {
	width := sized.GetWidth()
	sized.SetHeight(10)
	expectedArea := 10 * width
	acutalArea := sized.GetHeight() * sized.GetWidth()
	fmt.Println("Expected an area of ", expectedArea, " ,but got ", acutalArea)
}

func main() {
	rc := &Rectangle{
		height: 2,
		width:  3,
	}
	UseIt(rc)

	//Make UseIt wrong
	sq := NewSquare(5)
	UseIt(sq)

	//Approch to make test passed
	sq2 := Square2{5}
	rc2 := sq2.Rectangle()
	UseIt(&rc2)
}
