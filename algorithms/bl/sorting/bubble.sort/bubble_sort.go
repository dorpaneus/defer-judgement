package main

import (
	"fmt" // Import the fmt package for formatted input/output (like printing to the console)
)

// bubbleSort implements the Bubble Sort algorithm to sort an integer slice in ascending order.
// It takes a slice of integers 'arr' and modifies it in place.
func bubbleSort(arr []int) {
	// n is the length (number of elements) of the slice.
	n := len(arr) 

	// The outer loop iterates (n-1) times. A maximum of n-1 passes is required to sort n elements.
	// 'i' tracks how many elements at the end are already in their final sorted position.
	for i := 0; i < n-1; i++ { 
		// Flag to track if any swaps occurred in this pass. Used for optimization.
		swapped := false

		// The inner loop compares adjacent elements and swaps them if they are in the wrong order.
		// 'n-i-1' is used because the last 'i' elements are already sorted and can be skipped.
		for j := 0; j < n-i-1; j++ { 
			// If the current element is greater than the next element, they are out of order (for ascending sort).
			if arr[j] > arr[j+1] {
				// Swap arr[j] and arr[j+1] using Go's multiple assignment feature.
				arr[j], arr[j+1] = arr[j+1], arr[j]
				// Set the flag because a swap occurred.
				swapped = true
			}
		}

		// Optimization: If no two elements were swapped by the inner loop,
		// the array is already sorted, and we can break out of the outer loop early.
		if !swapped {
			break
		}
	}

}

// main is the entry point of the Go program.
func main() {
	// Initialize a slice of integers to be sorted.
	numbers := []int{4, 5, 3, 7, 1, 2, 0, -1, -3, 9, -6}
	
	// Print the slice before sorting.
	fmt.Println("Our original list of numbers is:", numbers)
	
	// Call the bubbleSort function. Since slices are passed by reference (header/pointer),
	// the function modifies the 'numbers' slice directly.
	bubbleSort(numbers) 
	
	// Print the slice after sorting to show the result.
	fmt.Println("Our sorted list of numbers is:", numbers)
}
