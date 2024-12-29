package main

import "fmt"

func heapify(arr []int, size, i int) {
	largest := i
	left := 2*i + 1
	right := 2*i + 2
	if left < size && arr[left] > arr[largest] {
		largest = left
	}

	if right < size && arr[right] > arr[largest] {
		largest = right
	}

	if largest != i {
		arr[largest], arr[i] = arr[i], arr[largest]
		heapify(arr, size, largest)
	}
}

func HeapSort(arr []int) {
	size := len(arr)
	for i := size/2 - 1; i >= 0; i-- {
		heapify(arr, size, i)
	}

	for i := size - 1; i > 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		heapify(arr, i, 0)
	}
}

func main() {
	arr := []int{1, 4, 29, 10}
	HeapSort(arr)
	fmt.Println(arr)

}
