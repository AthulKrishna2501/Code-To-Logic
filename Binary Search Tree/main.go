package main

import (
	"fmt"
	"math"
)

type Node struct {
	data  int
	left  *Node
	right *Node
}

type BST struct {
	root *Node
}

func (bst *BST) Insert(value int) {
	newNode := &Node{data: value}
	if bst.root == nil {
		bst.root = newNode
		return
	}

	bst.insertHelper(bst.root, newNode)
}

func (bst *BST) insertHelper(current, newNode *Node) {
	if newNode.data < current.data {
		if current.left == nil {
			current.left = newNode
		} else {
			bst.insertHelper(current.left, newNode)
		}
	} else if newNode.data > current.data {
		if current.right == nil {
			current.right = newNode
			return
		} else {
			bst.insertHelper(current.right, newNode)
		}
	}
}

func (bst *BST) Search(value int) bool {
	return bst.searchHelper(bst.root, value)
}

func (bst *BST) searchHelper(node *Node, value int) bool {
	if node == nil {
		return false
	}

	if node.data == value {
		return true
	}

	if value < node.data {
		return bst.searchHelper(node.left, value)
	}

	return bst.searchHelper(node.right, value)
}

func (bst *BST) Inorder() {
	bst.inorderHelper(bst.root)
	fmt.Println()
}

func (bst *BST) inorderHelper(node *Node) {
	if node == nil {
		return
	}

	bst.inorderHelper(node.left)
	fmt.Print(node.data, " ")
	bst.inorderHelper(node.right)
}

func (bst *BST) Preorder() {
	bst.preorderHelper(bst.root)
	fmt.Println()
}

func (bst *BST) preorderHelper(node *Node) {
	if node == nil {
		return
	}

	fmt.Print(node.data, " ")
	bst.preorderHelper(node.left)
	bst.preorderHelper(node.right)

}

func (bst *BST) PostOrder() {
	bst.postorderHelper(bst.root)
	fmt.Println()

}

func (bst *BST) postorderHelper(node *Node) {
	if node == nil {
		return
	}

	bst.postorderHelper(node.left)
	bst.postorderHelper(node.right)
	fmt.Print(node.data, " ")

}

func (bst *BST) LevelOrder() {
	bst.levelorderHelper(bst.root)
	fmt.Println()
}

func (bst *BST) levelorderHelper(node *Node) {
	if node == nil {
		return
	}

	queue := []*Node{bst.root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		fmt.Print(node.data, " ")
		if node.left != nil {
			queue = append(queue, node.left)
		}

		if node.right != nil {
			queue = append(queue, node.right)
		}
	}
}

func (bst *BST) Delete(node *Node, value int) *Node {
	if node == nil {
		return nil
	}

	if value < node.data {
		node.left = bst.Delete(node.left, value)
	} else if value > node.data {
		node.right = bst.Delete(node.right, value)
	} else {
		if node.left == nil && node.right == nil {
			return nil
		} else if node.left == nil {
			return node.right
		} else if node.right == nil {
			return node.left
		} else {
			successor := findMin(node.right)
			node.data = successor.data
			node.right = bst.Delete(node.right, successor.data)
		}
	}
	return node
}

func findMin(node *Node) *Node {
	current := node
	for current.left != nil {
		current = current.left
	}
	return current
}
func (bst *BST) FindClosestValue(target int) int {
	return bst.findClosestValue(bst.root, target, bst.root.data)
}

func (bst *BST) findClosestValue(node *Node, target, closest int) int {
	if node == nil {
		return closest

	}

	if math.Abs(float64(target)-float64(node.data)) < math.Abs(float64(target)-float64(closest)) {
		closest = node.data
	}

	if target < node.data {
		return bst.findClosestValue(node.left, target, closest)
	}
	if target > node.data {
		return bst.findClosestValue(node.right, target, closest)
	}
	return closest
}

func isBST(node *Node) bool {
	return validateBST(node, math.MinInt, math.MaxInt)
}

func validateBST(node *Node, min, max int) bool {
	if node == nil {
		return true
	}

	if node.data <= min || node.data >= max {
		return false
	}

	return validateBST(node.left, min, node.data) && validateBST(node.right, node.data, max)
}

func main() {
	root := &BST{}
	root.Insert(20)
	root.Insert(10)
	root.Insert(30)
	root.Insert(5)
	root.Insert(15)
	root.Insert(25)
	root.Insert(35)
	root.Inorder()
	root.Preorder()
	root.PostOrder()
	root.LevelOrder()
	target := 13
	fmt.Println(root.FindClosestValue(target))
	if isBST(&Node{}) {
		fmt.Println("isBSt")
	} else {
		fmt.Println("NotBST")
	}

	root.Delete(root.root, 20)

}
