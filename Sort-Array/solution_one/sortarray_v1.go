package main

import "fmt"

func main() {

	arr := []int{1, 4, 5, 2, 5, 3, 2, 2, 5, 7, 6}
	sort := SortNumbers(arr)

	arr2 := []int{}
	sort2 := SortNumbers(arr2)

	fmt.Println(sort)
	fmt.Println(sort2)

}

func SortNumbers(arr []int) []int {
	if arr == nil {
		return nil
	} else {
		for i := 0; i <= len(arr)-1; i++ {
			for j := 0; j < len(arr)-1-i; j++ {
				if arr[j] > arr[j+1] {
					arr[j], arr[j+1] = arr[j+1], arr[j]
				}
			}
		}
	}
	return arr
}
