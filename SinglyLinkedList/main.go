package main

import "fmt"

type Node struct {
	value int
	next  *Node
}

type LinkedList struct {
	head *Node
}

func (list *LinkedList) displayLinkedList() {
	current := list.head
	for current != nil {
		fmt.Println(current.value)
		current = current.next
	}
}

func (list *LinkedList) insertAtFirst(value int) {
	newNode := &Node{value: value}
	if list.head == nil {
		list.head = newNode
		return
	}

	newNode.next = list.head
	list.head = newNode
}

func (list *LinkedList) insertAtEnd(value int) {
	newNode := &Node{value: value}
	if list.head == nil {
		list.head = newNode
		return
	}

	current := list.head
	for current.next != nil {
		current = current.next
	}

	current.next = newNode
}

func (list *LinkedList) insertAfterElement(value, pos int) {
	newNode := &Node{value: value}
	if list.head.value == pos {
		newNode.next = list.head.next
		list.head.next = newNode
		return
	}

	current := list.head
	for current != nil && current.value != pos {
		current = current.next
	}
	newNode.next = current.next
	current.next = newNode
}

func (list *LinkedList) insertBeforeElement(value, x int) {
	newNode := &Node{value: value}
	if list.head.value == x {
		newNode.next = list.head
		list.head = newNode
		return
	}

	current := list.head

	for current.next != nil && current.next.value != x {
		current = current.next
	}

	newNode.next = current.next
	current.next = newNode

}

func (list *LinkedList) insertAtNth(value, x int) {
	newNode := &Node{value: value}
	if list.head.value == x {
		list.head = newNode
		return
	}

	current := list.head
	for current != nil && current.value != x {
		current = current.next
	}

	current.value = newNode.value
	newNode.next = current.next
}

func (list *LinkedList) deleteFirst() {
	list.head = list.head.next
}

func (list *LinkedList) deleteEnd() {
	if list.head.next == nil {
		list.head = nil
		return
	}

	current := list.head
	for current.next.next != nil {
		current = current.next
	}

	current.next = nil
}

func (list *LinkedList) deleteAfterElement(pos int) {
	if list.head.value == pos {
		list.head.next = list.head.next.next
		return
	}

	current := list.head
	for current != nil && current.value != pos {
		current = current.next
	}

	current.next = current.next.next
}

func (list *LinkedList) deleteBeforeElement(pos int) {
	var prev *Node
	current := list.head
	for current.next != nil && current.next.value != pos {
		prev = current
		current = current.next
	}
	prev.next = current.next
}

func main() {
	list := &LinkedList{}
	list.insertAtFirst(1)
	list.insertAtFirst(2)
	list.insertAtFirst(3)
	list.insertAtEnd(10)
	list.insertAfterElement(30, 2)
	list.insertBeforeElement(9, 10)
	list.insertAtNth(8, 10)
	list.deleteFirst()
	list.deleteEnd()
	list.deleteAfterElement(1)
	list.deleteBeforeElement(1)
	list.displayLinkedList()
}
