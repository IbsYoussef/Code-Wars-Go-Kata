package main

import "fmt"

func main() {

	// var a = [5]int{11, -4, 7, 8, -10}
	// min, max := findMinAndMax(a)
	// fmt.Println("Min: ", min)
	// fmt.Println("Max: ", max)

	array := []int{1, 2, 3, 4}
	a := MinMax(array)
	fmt.Print(a)

	arr := []int{2, 256, 500}
	b := MinMax(arr)
	fmt.Print(b)

	brr := []int{55, 66, 100, 96}
	c := MinMax(brr)
	fmt.Print(c)

}

// func findMinAndMax(a [5]int) (min int, max int) {
// 	min = a[0]
// 	max = a[0]
// 	for _, value := range a {
// 		if value < min {
// 			min = value
// 		}
// 		if value > max {
// 			max = value
// 		}
// 	}
// 	return min, max
// }

// func main() {
// 	array := []int{1, 2, 3, 4, 5}
// 	fmt.Print(MinMax(array))
// }

func MinMax(a []int) [2]int {
	min := a[0]
	max := a[0]
	list := [2]int{}
	for _, value := range a {
		if value < min {
			min = value
		}
		if value > max {
			max = value
		}
	}
	list[0] = min
	list[1] = max
	return list
}
