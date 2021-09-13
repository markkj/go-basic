package main

import "fmt"

type Node struct {
	value               int
	left, right, parent *Node
}

func NewNode(value int, left, right *Node) *Node {
	n := &Node{value: value, left: left, right: right}
	left.parent = n
	right.parent = n
	return n
}

type BinaryTree struct {
	root *Node
}

func NewBinaryTree(root *Node) *BinaryTree {
	return &BinaryTree{root: root}
}

type InOrderIterator struct {
	Current       *Node
	root          *Node
	returnedStart bool
}

func NewInOrderIterator(root *Node) *InOrderIterator {
	i := &InOrderIterator{root: root, Current: root, returnedStart: false}
	for i.Current.left != nil {
		i.Current = i.Current.left
	}
	return i
}

func (i *InOrderIterator) Resest() {
	i.Current = i.root
	i.returnedStart = false
}

func (i *InOrderIterator) MoveNext() bool {
	if i.Current == nil {
		return false
	}
	if !i.returnedStart {
		i.returnedStart = true
		return true
	}

	if i.Current.right != nil {
		i.Current = i.Current.right
		for i.Current.left != nil {
			i.Current = i.Current.left
		}
		return true
	} else {
		p := i.Current.parent
		for p != nil && i.Current == p.right {
			i.Current = p
			p = p.parent
		}
		i.Current = p
		return i.Current != nil
	}

}

func (b *BinaryTree) InOrder() *InOrderIterator {
	return NewInOrderIterator(b.root)
}

func main() {
	root := NewNode(10, &Node{value: 20}, &Node{value: 30})

	it := NewInOrderIterator(root)
	for it.MoveNext() {
		fmt.Printf("%d,", it.Current.value)
	}

	bt := NewBinaryTree(root)
	it = bt.InOrder()
	for it.MoveNext() {
		fmt.Printf("%d,", it.Current.value)
	}
}
