package main

import "fmt"

type MinHeap struct {
	array []int
}

func (h *MinHeap) Insert(value int) {
	h.array = append(h.array, value)
	h.heapifyUp(len(h.array) - 1)
}

func (h *MinHeap) heapifyUp(index int) {
	for index > 0 {
		parent := (index - 1) / 2
		if h.array[index] < h.array[parent] {
			h.array[index], h.array[parent] = h.array[parent], h.array[index]
			index = parent
		} else {
			break
		}
	}
}

func (h *MinHeap) Remove() {
	h.array[0] = h.array[len(h.array)-1]
	h.array = h.array[:len(h.array)-1]
	h.heapifyDown(0)
}

func (h *MinHeap) heapifyDown(index int) {
	size := len(h.array)
	smallest := index
	left := 2*index + 1
	right := 2*index + 2
	if left < size && h.array[left] < h.array[smallest] {
		smallest = left
	}

	if right < size && h.array[right] < h.array[smallest] {
		smallest = right
	}

	if smallest != index {
		h.array[index], h.array[smallest] = h.array[smallest], h.array[index]
		h.heapifyDown(smallest)
	}
}

func (h *MinHeap) Display() {
	fmt.Print(h.array)
}

func main() {
	heap := &MinHeap{}
	heap.Insert(20)
	heap.Insert(10)
	heap.Insert(30)
	heap.Insert(50)
	heap.Remove()
	heap.Display()
}
