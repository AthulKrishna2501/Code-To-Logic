package main

import "fmt"

func findIndex(arr []int, x int) int {
	left, right := 0, len(arr)-1
	index := 0
	for left <= right {
		mid := (left + right) / 2
		if arr[mid] == x {
			index = mid
		} else if arr[mid] < x {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return index
}

func main() {
	arr := []int{1, 2, 2, 3, 4, 5}
	x := 2
	fmt.Println(findIndex(arr, x))
}
