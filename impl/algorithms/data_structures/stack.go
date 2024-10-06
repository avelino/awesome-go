package main

import "fmt"

type Node struct {
    data int
    next *Node
}

func newNode(data int) *Node {
    return &Node{data: data, next: nil}
}

func (n *Node) display() {
    fmt.Println(n.data)
}

type Stack struct{}

func (s *Stack) push(head *Node, value int) *Node {
    new_node := newNode(value)
    new_node.next = head
    head = new_node
    return head
}

func (s *Stack) pop(head **Node) int {
    if *head == nil {
        fmt.Println("Stack is empty")
        return -1
    }
    value := (*head).data
    temp := *head
    *head = (*head).next
    temp.next = nil
    return value
}

func (s *Stack) peek(head *Node) int {
    if head == nil {
        fmt.Println("Stack is empty")
        return -1
    }
    return head.data
}

func (s *Stack) size(head *Node) int {
    size := 0
    temp := head
    for temp != nil {
        size++
        temp = temp.next
    }
    return size
}

func (s *Stack) isEmpty(head *Node) bool {
    return head == nil
}

func main() {
    var head *Node = nil
    var s Stack
    head = s.push(head, 5)
    head = s.push(head, 10)
    head = s.push(head, 15)
    fmt.Println("Top element: ", s.peek(head))
    fmt.Println("Stack size: ", s.size(head))
    for !s.isEmpty(head) {
        fmt.Print(s.pop(&head), " ")
    }
}
