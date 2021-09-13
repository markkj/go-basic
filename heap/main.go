package main

import "fmt"

type HeapInterface interface {
	push(val int32)
	pop() int32
	heapifyUp(index int)
	heapifyDown(index int)
	swap(i1 int, i2 int)
	peek() int32
	len() int
}

type minHeap struct {
	arr []int32
}

func main() {
	minheap := NewMinHeap([]int32{7, 6, 5, 4, 3, 2, 1})
	// minheap.push(2)
	fmt.Println(minheap)
	fmt.Println(minheap.peek())
	fmt.Println(minheap.pop())
	fmt.Println(minheap.pop())

}

func (m *minHeap) push(val int32) {
	m.arr = append(m.arr, val)
	m.heapifyUp(len(m.arr) - 1)
}

func (m *minHeap) pop() int32 {
	if len(m.arr) == 0 {
		return -1
	}

	p := m.arr[0]
	l := len(m.arr) - 1
	m.arr[0] = m.arr[l]
	m.arr = m.arr[:l]

	m.heapifyDown(0)
	return p
}

func (m *minHeap) peek() int32 {
	return m.arr[0]
}

func (m *minHeap) len() int {
	return len(m.arr)
}

func (m *minHeap) heapifyUp(index int) {
	for m.arr[parent(index)] > m.arr[index] {
		m.swap(parent(index), index)
		index = parent(index)
	}
}

func (m *minHeap) heapifyDown(index int) {
	lastIndex := len(m.arr) - 1
	l, r := left(index), right(index)

	childCompare := 0

	for l <= lastIndex {
		if l == lastIndex {
			childCompare = l
		} else if m.arr[l] < m.arr[r] {
			childCompare = l
		} else {
			childCompare = r
		}

		if m.arr[index] > m.arr[childCompare] {
			m.swap(index, childCompare)
			index = childCompare
			l, r = left(index), right(index)
		} else {
			return
		}
	}
}

func parent(index int) int {
	return (index - 1) / 2
}

func left(index int) int {
	return index*2 + 1

}

func right(index int) int {
	return index*2 + 2
}

func (m *minHeap) swap(i1 int, i2 int) {
	m.arr[i1], m.arr[i2] = m.arr[i2], m.arr[i1]
}

func NewMinHeap(arr []int32) HeapInterface {
	m := &minHeap{}
	for _, v := range arr {
		m.push(v)
	}
	return m
}
