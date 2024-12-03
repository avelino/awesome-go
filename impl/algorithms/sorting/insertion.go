package main

import "fmt"

// Insertion sort function
func insertionSort(arr []int) {
	// Iterate through each element of the array
	for i := 1; i < len(arr); i++ {
		// Save the current element and its index
		key := arr[i]
		j := i - 1

		// Shift all elements that are greater than the current element
		// one position to the right
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}

		// Insert the current element in its correct position
		arr[j+1] = key
	}
}

func main() {
	// Example usage
	arr := []int{5, 2, 4, 6, 1, 3}
	fmt.Println("Before sorting:", arr)
	insertionSort(arr)
	fmt.Println("After sorting:", arr)
}
