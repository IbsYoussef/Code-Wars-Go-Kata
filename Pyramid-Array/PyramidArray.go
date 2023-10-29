package main

import "fmt"

func main() {

	fmt.Println(Pyramid(2))
	fmt.Println(Pyramid(4))
	fmt.Println(Pyramid(6))

}

func Pyramid(n int) [][]int {
	pyramid := make([][]int, n)

	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			pyramid[i] = append(pyramid[i], 1)
		}
	}

	return pyramid
}
