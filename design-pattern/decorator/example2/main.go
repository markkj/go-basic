package main

import "fmt"

type Shape interface {
	Render() string
}

type Circle struct {
	Radius float32
}

func (c *Circle) Render() string {
	return fmt.Sprintf("Circle of radius %f", c.Radius)
}

func (c *Circle) Resize(factor float32) {
	c.Radius *= factor
}

type Square struct {
	Side float32
}

func (c *Square) Render() string {
	return fmt.Sprintf("Square with side %f", c.Side)
}

type ColoredShape struct {
	Shape
	Color string
}

func (c *ColoredShape) Render() string {
	return fmt.Sprintf("%s with color %s", c.Shape.Render(), c.Color)
}

type TransparentSahpe struct {
	Shape
	Transparency float32
}

func (c *TransparentSahpe) Render() string {
	return fmt.Sprintf("%s with trasparency %f", c.Shape.Render(), c.Transparency)
}

func main() {
	c := Circle{2}
	c.Resize(10)
	fmt.Println(c.Render())

	redCircle := ColoredShape{&c, "Red"}

	fmt.Println(redCircle.Render())

	traCircle := TransparentSahpe{&c, 0.3}
	fmt.Println(traCircle.Render())

}
