package main

import (
	"fmt"
)

func selectionSort(arr []int) {
	// Loop through each element in the array
	for i := 0; i < len(arr)-1; i++ {
		// Assume the current element is the minimum
		min := i

		// Loop through the remaining elements in the array
		for j := i + 1; j < len(arr); j++ {
			// If the current element is less than the assumed minimum, update the minimum
			if arr[j] < arr[min] {
				min = j
			}
		}

		// Swap the current element with the minimum element
		arr[i], arr[min] = arr[min], arr[i]
	}
}

func main() {
	// Example usage
	arr := []int{5, 2, 4, 6, 1, 3}
	fmt.Println("Before sorting:", arr)
	selectionSort(arr)
	fmt.Println("After sorting:", arr)
}
