package main

import "fmt"

//Low-Level
type Keyboard struct{}
type Monitor struct{}
type RBGKeyboard struct{}

//High-Level
type Computer struct {
	// Keyboard //Can change Keyboard
	// Monitor  //Can change Monitor
	keyboard KeyboardInterface
	monitor  MonitorInterface
}

type KeyboardInterface interface {
	typing()
}

type MonitorInterface interface {
	display()
}

func (k *Keyboard) typing() {
	fmt.Println("Typing with normal keyboard")
}

func (k *RBGKeyboard) typing() {
	fmt.Println("Typing with RGB keyboard")
}

func (m *Monitor) display() {
	fmt.Println("Display with normal monitor")
}

func main() {
	c := &Computer{
		keyboard: &Keyboard{},
		monitor:  &Monitor{},
	}
	c.keyboard.typing()
	c.monitor.display()

	c1 := &Computer{
		keyboard: &RBGKeyboard{},
		monitor:  &Monitor{},
	}
	c1.keyboard.typing()
	c1.monitor.display()
}
