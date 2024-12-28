package main

import "fmt"

// MinHeap
type MinHeap struct {
	arr []int
}

func (h *MinHeap) Insert(value int) {
	h.arr = append(h.arr, value)
	h.heapifyUp(len(h.arr) - 1)
}

func (h *MinHeap) heapifyUp(index int) {
	for index > 0 {
		parent := (index - 1) / 2
		if h.arr[index] < h.arr[parent] {
			h.arr[parent], h.arr[index] = h.arr[index], h.arr[parent]
			index = parent
		} else {
			break
		}
	}

}

func (h *MinHeap) Remove() int {
	root := h.arr[0]
	h.arr = h.arr[:len(h.arr)-1]
	h.heapifyDown(0)
	return root
}

func (h *MinHeap) heapifyDown(index int) {
	size := len(h.arr)
	smallest := index
	left := 2*index + 1
	right := 2*index + 2

	if left < size && h.arr[left] < h.arr[smallest] {
		smallest = left
		
	}

	if right < size && h.arr[right] < h.arr[smallest] {
		smallest = right
	}

	if smallest != index {
		h.arr[index], h.arr[smallest] = h.arr[smallest], h.arr[index]
		h.heapifyDown(smallest)
	}
}

func (h *MinHeap) Display() {
	fmt.Println(h.arr)
}

func main() {
	h := &MinHeap{}
	h.Insert(10)
	h.Insert(20)
	h.Insert(5)
	h.Insert(30)
	h.Insert(15)
	h.Display()
	h.Remove()
	h.Display()

}
