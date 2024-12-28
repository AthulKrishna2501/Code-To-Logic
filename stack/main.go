package main

import (
	"fmt"
	"math"
)

type Stack struct {
	items []int
}

func (s *Stack) push(value int) {
	s.items = append(s.items, value)
}

func (s *Stack) pop() int {
	top := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return top
}

func (s *Stack) isEmpty() bool {
	return len(s.items) == 0
}

func findSmallest(stack *Stack) int {
	min := math.MaxInt
	tempStack := &Stack{}

	for !stack.isEmpty() {
		value := stack.pop()
		if value < min {
			min = value
		}
		tempStack.push(value)
	}

	for !tempStack.isEmpty() {
		stack.push(tempStack.pop())
	}
	return min
}

func removeMiddle(stack *Stack) {
	mid := len(stack.items) / 2
	tempStack := &Stack{}
	for i := 0; i < mid; i++ {
		tempStack.push(stack.pop())
	}
	stack.pop()

	for !tempStack.isEmpty() {
		stack.push(tempStack.pop())
	}
}

func main() {
	stack := &Stack{}
	stack.push(10)
	stack.push(2)
	stack.push(4)
	stack.push(8)
	stack.push(0)
	fmt.Println(stack)
	removeMiddle(stack)
	fmt.Println(stack)
	fmt.Println(findSmallest(stack))
}


