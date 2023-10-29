package main

import (
	"fmt"
)

func main() {

	fmt.Println(SummationV1(5))
	fmt.Println(SummationV2(5))

}

func SummationV1(n int) int {
	// the sleeper must awaken!
	var result int = (n * (n + 1)) / 2
	return result
}

func SummationV2(n int) int {
	if n == 1 {
		return 1
	} else {
		return n + SummationV2(n-1)
	}
}
