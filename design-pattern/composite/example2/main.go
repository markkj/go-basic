package main

import "fmt"

type NeuronInterface interface {
	Iter() []*Neuron
}

type Neuron struct {
	In, Out []*Neuron
}

func (n *Neuron) ConnectTo(other *Neuron) {
	n.Out = append(n.Out, other)
	other.In = append(other.In, n)
}
func (n *Neuron) Iter() []*Neuron {
	return []*Neuron{n}
}

type NeuronLayer struct {
	Neuron []Neuron
}

func (n *NeuronLayer) Iter() []*Neuron {
	result := make([]*Neuron, 0)
	for i := range n.Neuron {
		result = append(result, &n.Neuron[i])
	}
	return result
}

func NewLayer(count int) *NeuronLayer {
	return &NeuronLayer{make([]Neuron, count)}
}

func Connect(left, right NeuronInterface) {
	for _, l := range left.Iter() {
		for _, r := range right.Iter() {
			l.ConnectTo(r)
		}
	}
}

func main() {
	n1, n2 := &Neuron{}, &Neuron{}
	layer1, layer2 := NewLayer(3), NewLayer(2)

	Connect(n1, n2)
	fmt.Printf("%p", n1)
	fmt.Println(n1)
	fmt.Println(n2)

	Connect(n1, layer1)
	Connect(layer2, n1)
	Connect(layer1, layer2)
}
