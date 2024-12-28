package main

import "fmt"

func binarySearch(target int, arr []int) int {
	left, right := 0, len(arr)-1
	mid := (right-left)/2 + left

	for left <= right {
		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			left = mid + 1
			
		} else {
			right = mid - 1
		}
	}
	return -1
}
func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	target := 2
	result := binarySearch(target, arr)
	fmt.Println("Targest occured:", result)
}
