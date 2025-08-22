package main

import "fmt"

type Node struct {
	data int
	next *Node
}

func newNode(data int) *Node {
	return &Node{data: data, next: nil}
}

func (n *Node) size(head *Node) int {
	if head == nil {
		fmt.Println("Linkedlist is empty")
		return 0
	}

	size := 0
	for head != nil {
		size++
		head = head.next
	}

	return size
}

func (n *Node) display(head *Node) {
	if head == nil {
		fmt.Println("Linkedlist is empty")
		return
	}

	for head != nil {
		fmt.Print(head.data, " ")
		head = head.next
	}
	fmt.Println()
}

func (n *Node) insertAtHead(data int, head *Node) *Node {
	newNode := newNode(data)
	newNode.next = head
	return newNode
}

func (n *Node) insertAtTail(data int, head *Node) *Node {
	newNode := newNode(data)
	if head == nil {
		return newNode
	}

	current := head
	for current.next != nil {
		current = current.next
	}

	current.next = newNode
	return head
}

func (n *Node) insertAtPosition(headRef **Node, position, data int) *Node {
	newNode := newNode(data)

	if *headRef == nil || position == 0 {
		newNode.next = *headRef
		*headRef = newNode
		return *headRef
	}

	current := *headRef
	for i := 0; current != nil && i < position-1; i++ {
		current = current.next
	}

	if current == nil {
		return *headRef
	}

	newNode.next = current.next
	current.next = newNode

	return *headRef
}

func (n *Node) deleteAtHead(head *Node) *Node {
	if head == nil {
		return nil
	}
	return head.next
}

func (n *Node) deleteAtTail(head *Node) *Node {
	if head == nil {
		fmt.Println("Linked list is empty.")
		return nil
	}

	if head.next == nil {
		return nil
	}

	cur := head
	for cur.next.next != nil {
		cur = cur.next
	}

	cur.next = nil
	return head
}

func (n *Node) deleteAtPosition(position int, head *Node) *Node {
	if head == nil {
		fmt.Println("Linked list is empty.")
		return nil
	}

	if position == 0 {
		return head.next
	}

	current := head
	var prev *Node

	for current != nil && position > 0 {
		prev = current
		current = current.next
		position--
	}

	if current == nil {
		fmt.Println("Invalid position")
		return head
	}

	prev.next = current.next
	return head
}

func main() {
	// create a linked list and perform operations
	var head *Node
	head = newNode(1)
	head = head.insertAtTail(2, head)
	head = head.insertAtTail(3, head)
	head = head.insertAtHead(0, head)

	head.display(head)           // expected output: 0 1 2 3
	fmt.Println(head.size(head)) // expected output: 4

	head = head.deleteAtHead(head)
	head.display(head) // expected output: 1 2 3

	head = head.deleteAtTail(head)
	head.display(head) // expected output: 1 2

	head = head.deleteAtPosition(1, head)
	head.display(head) // expected output: 1

	head = head.deleteAtPosition(0, head)
	head.display(head) // expected output:

	head = head.insertAtTail(2, head)
	head = head.insertAtTail(3, head)
	head = head.insertAtHead(1, head)
	head = head.insertAtPosition(&head, 1, 4)
	head.display(head) // expected output: 1 4 2 3
}
