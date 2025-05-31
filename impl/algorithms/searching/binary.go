package main

import "fmt"

// binarySearch returns the index of key in the given sorted slice of integers.
// If key is not found, it returns -1.
func binarySearch(arr []int, key int) int {
	// Initialize low and high indices
	low := 0
	high := len(arr) - 1

	// Loop until low index is less than or equal to high index
	for low <= high {
		// Calculate middle index
		mid := (low + high) / 2

		// If key is found at middle index, return it
		if arr[mid] == key {
			return mid
		}

		// If key is less than the middle element, search in the left half
		if arr[mid] > key {
			high = mid - 1
		}

		// If key is greater than the middle element, search in the right half
		if arr[mid] < key {
			low = mid + 1
		}
	}

	// If key is not found, return -1
	return -1
}

func main() {
	// Example usage of binarySearch function
	arr := []int{2, 3, 4, 10, 40}
	key := 10
	index := binarySearch(arr, key)
	if index == -1 {
		fmt.Println("Element not present")
	} else {
		fmt.Printf("Element found at index %d\n", index)
	}
}
