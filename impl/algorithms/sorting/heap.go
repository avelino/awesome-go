package main

import "fmt"

// The heapify function takes a slice, a length n and an index i,
// and ensures that the subtree rooted at index i satisfies the heap property.
// The heap property is satisfied if the parent node at index i is larger than or equal to
// its children.
func heapify(arr []int, n, i int) {
    largest := i       // Initialize largest as root
    left := 2*i + 1    // left child of i
    right := 2*i + 2   // right child of i

    // If left child is larger than root
    if left < n && arr[left] > arr[largest] {
        largest = left
    }

    // If right child is larger than largest so far
    if right < n && arr[right] > arr[largest] {
        largest = right
    }

    // If largest is not root
    if largest != i {
        arr[i], arr[largest] = arr[largest], arr[i] // swap
        // Recursively heapify the affected sub-tree
        heapify(arr, n, largest)
    }
}

// The heapSort function takes a slice and sorts it in increasing order
func heapSort(arr []int) {
    n := len(arr)

    // Build heap (rearrange array)
    // Starting with the last non-leaf node and heapifying all nodes in reverse level order
    for i := n/2 - 1; i >= 0; i-- {
        heapify(arr, n, i)
    }

    // One by one extract an element from heap
    for i := n - 1; i > 0; i-- {
        // Move current root to end
        arr[0], arr[i] = arr[i], arr[0]

        // call max heapify on the reduced heap
        heapify(arr, i, 0)
    }
}

func main() {
    // Example usage
    arr := []int{12, 11, 13, 5, 6, 7}
    fmt.Println("Original array:", arr)
    heapSort(arr)
    fmt.Println("Sorted array:", arr)
}
