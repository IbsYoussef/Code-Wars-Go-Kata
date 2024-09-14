package main

import (
	"fmt"
	"sort"
)

func main() {
	// Test case 1: Random numbers
	nums1 := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}
	SortNumbers(nums1)
	fmt.Printf("Test case 1 sorted: %v\n", nums1)

	// Test case 2: Already sorted
	nums2 := []int{1, 2, 3, 4, 5}
	SortNumbers(nums2)
	fmt.Printf("Test case 2 sorted: %v\n", nums2)

	// Test case 3: Reverse sorted
	nums3 := []int{9, 8, 7, 6, 5}
	SortNumbers(nums3)
	fmt.Printf("Test case 3 sorted: %v\n", nums3)

	// Test case 4: All the same numbers
	nums4 := []int{5, 5, 5, 5, 5}
	SortNumbers(nums4)
	fmt.Printf("Test case 4 sorted: %v\n", nums4)

	// Test case 5: Negative numbers
	nums5 := []int{-1, -3, -2, -4, -5}
	SortNumbers(nums5)
	fmt.Printf("Test case 5 sorted: %v\n", nums5)

	// Test case 6: Mixture of positive and negative numbers
	nums6 := []int{-1, 3, -2, 4, -5}
	SortNumbers(nums6)
	fmt.Printf("Test case 6 sorted: %v\n", nums6)

	// Test case 7: Empty array
	nums7 := []int{}
	SortNumbers(nums7)
	fmt.Printf("Test case 7 sorted: %v\n", nums7)

	// Test case 8: Single element
	nums8 := []int{42}
	SortNumbers(nums8)
	fmt.Printf("Test case 8 sorted: %v\n", nums8)
}

func SortNumbers(numbers []int) []int {
	if len(numbers) == 0 {
		return numbers
	}
	sort.Ints(numbers)
	return numbers
}
