package main

import "fmt"

const QUEUE_SIZE int = 10

var queue [QUEUE_SIZE]int
var rear int = -1
var front int = -1

func isEmpty() bool {
return front == -1
}

func peek() int {
if isEmpty() {
fmt.Println("Queue is empty")
return -1
}
return queue[front]
}

func isFull() bool {
    return rear >= QUEUE_SIZE-1
    }
    
    func enqueue(data int) {
    if isFull() {
    fmt.Println("Queue Overflow")
    return
    }rear += 1
    queue[rear] = data
    
    if front == -1 {
        front = 0
    }
    }
    
    func dequeue() {
    if isEmpty() {
    fmt.Println("Queue Underflow")
    return
    }
    
    front += 1

    if front > rear {
        front = -1
        rear = -1
    }
    }
    
    func size() int {
    if isEmpty() {
    return 0
    }
    
    return rear - front + 1
}

func display() {
if isEmpty() {
fmt.Println("Queue is empty")
return
}

for i := front; i <= rear; i++ {
	fmt.Printf("%d ", queue[i])
}
fmt.Println()
}

func main() {
// testing the queue implementation
enqueue(1)
enqueue(2)
enqueue(3)
enqueue(4)
display() // output: 1 2 3 4
dequeue()
dequeue()
display() // output: 3 4
enqueue(5)
enqueue(6)
enqueue(7)
enqueue(8)
enqueue(9)
enqueue(10)
display() // output: 3 4 5 6 7 8 9 10
enqueue(11) // output: Queue Overflow
dequeue()
dequeue()
dequeue()
dequeue()
dequeue()
dequeue()
dequeue()
dequeue()
dequeue()
display() // output: Queue is empty
dequeue() // output: Queue Underflow
}