package main

import (
	"fmt"
)

type Node struct {
	data  int
	left  *Node
	right *Node
}

type BinaryTree struct {
	root *Node
}

func NewBinaryTree() *BinaryTree {
	return &BinaryTree{}
}

func (bt *BinaryTree) Insert(data int) {
	newNode := &Node{data: data, left: nil, right: nil}

	if bt.root == nil {
		bt.root = newNode
	} else {
		queue := []*Node{bt.root}

		for len(queue) > 0 {
			node := queue[0]
			queue = queue[1:]

			if node.left == nil {
				node.left = newNode
				break
			} else {
				queue = append(queue, node.left)
			}

			if node.right == nil {
				node.right = newNode
				break
			} else {
				queue = append(queue, node.right)
			}
		}
	}
}

func (bt *BinaryTree) PreorderTraversal(node *Node) {
	if node != nil {
		fmt.Printf("%d ", node.data)
		bt.PreorderTraversal(node.left)
		bt.PreorderTraversal(node.right)
	}
}

func (bt *BinaryTree) InorderTraversal(node *Node) {
	if node != nil {
		bt.InorderTraversal(node.left)
		fmt.Printf("%d ", node.data)
		bt.InorderTraversal(node.right)
	}
}

func (bt *BinaryTree) PostorderTraversal(node *Node) {
	if node != nil {
		bt.PostorderTraversal(node.left)
		bt.PostorderTraversal(node.right)
		fmt.Printf("%d ", node.data)
	}
}

func main() {
	binaryTree := NewBinaryTree()

	binaryTree.Insert(1)
	binaryTree.Insert(2)
	binaryTree.Insert(3)
	binaryTree.Insert(4)
	binaryTree.Insert(5)

	fmt.Println("Preorder Traversal:")
	binaryTree.PreorderTraversal(binaryTree.root)
	fmt.Println()

	fmt.Println("Inorder Traversal:")
	binaryTree.InorderTraversal(binaryTree.root)
	fmt.Println()

	fmt.Println("Postorder Traversal:")
	binaryTree.PostorderTraversal(binaryTree.root)
	fmt.Println()
}