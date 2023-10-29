package main

import (
	"fmt"
)

func main() {
	fmt.Println(DigitalRoot(12))               // 3
	fmt.Println(DigitalRoot(3331))             // 1
	fmt.Println(DigitalRoot(54647))            // 8
	fmt.Println(DigitalRoot(5464785))          // 12
	fmt.Println(DigitalRoot(999999999999))     // 9
	fmt.Println(DigitalRoot(3456345634563456)) // 9
}

// My Solution, Needs work, Works except 1 random test not sure why, will revisit
// func DigitalRoot(n int) int {
// 	strarray := strings.Split(strconv.Itoa(n), "")
// 	Digitarr := []int{}

// 	for _, v := range strarray {
// 		j, err := strconv.Atoi(v)
// 		if err != nil {
// 			panic(err)
// 		}
// 		Digitarr = append(Digitarr, j)
// 	}

// 	root := 0
// 	for _, v := range Digitarr {
// 		root += v
// 	}

// 	for root > 9 {
// 		strarray := strings.Split(strconv.Itoa(root), "")
// 		Digitarr := []int{}

// 		for _, v := range strarray {
// 			j, err := strconv.Atoi(v)
// 			if err != nil {
// 				panic(err)
// 			}
// 			Digitarr = append(Digitarr, j)
// 		}

// 		root := 0
// 		for _, v := range Digitarr {
// 			root += v
// 		}
// 		return root
// 	}

// 	return root
// }

// Final Solution Found
func DigitalRoot(n int) int {
	if n < 10 {
		return n
	}
	n = DigitalRoot(n/10) + n%10
	return DigitalRoot(n)
}
