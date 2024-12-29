package main

import "fmt"

type Node struct {
	value int
	left  *Node
	right *Node
}

type Tree struct {
	root *Node
}

func (t *Tree) Inorder() {
	t.InorderHelper(t.root)
	fmt.Println()
}

func (t *Tree) InorderHelper(node *Node) {
	if node == nil {
		return
	}

	t.InorderHelper(node.left)
	fmt.Print(node.value, " ")
	t.InorderHelper(node.right)
}

func (t *Tree) Preorder() {
	t.preorderHelper(t.root)
	fmt.Println()
}

func (t *Tree) preorderHelper(node *Node) {
	if node == nil {
		return
	}

	fmt.Print(node.value, " ")
	t.preorderHelper(node.left)
	t.preorderHelper(node.right)
}

func (t *Tree) PostOrder() {
	t.postorderHelper(t.root)
	fmt.Println()
}

func (t *Tree) postorderHelper(node *Node) {
	if node == nil {
		return
	}

	t.postorderHelper(node.left)
	t.postorderHelper(node.right)
	fmt.Print(node.value, " ")
}

func (t *Tree) LevelOrder() {
	t.levelOrderHelper(t.root)
	fmt.Println()
}

func (t *Tree) levelOrderHelper(node *Node) {
	if node == nil {
		return
	}

	queue := []*Node{t.root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		fmt.Print(node.value, " ")
		if node.left != nil {
			queue = append(queue, node.left)
		}
		if node.right != nil {
			queue = append(queue, node.right)
		}
	}
	fmt.Println()

}

func main() {
	root := &Node{value: 1}
	root.left = &Node{value: 2}
	root.right = &Node{value: 3}
	root.left.left = &Node{value: 4}
	root.left.right = &Node{value: 5}
	root.right.right = &Node{value: 6}
	root.right.left = &Node{value: 7}

	tree := &Tree{root: root}
	fmt.Println("Inorder Traversal:")
	tree.Inorder()
}
