package main

import (
	"fmt"
)

type Driven interface {
	Drive()
}

type Car struct{}

func (c *Car) Drive() {
	fmt.Println("Car is being driven")
}

type Driver struct {
	Age int
}

type CarProxy struct {
	car    Car
	driver *Driver
}

func (c *CarProxy) Drive() {
	if c.driver.Age >= 18 {
		c.car.Drive()
	} else {
		fmt.Println("Driver too young !!")
	}
}

func NewCarProxy(driver *Driver) *CarProxy {
	return &CarProxy{Car{}, driver}
}

func main() {
	d := Driver{16}
	c := NewCarProxy(&d)
	c.Drive()

	d = Driver{20}
	c = NewCarProxy(&d)
	c.Drive()

}
