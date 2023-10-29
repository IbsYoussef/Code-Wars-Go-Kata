package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(TowerBuilder(6))
}

func TowerBuilder(nFloors int) []string {
	var res []string
	var Tower strings.Builder

	for i := 1; i < nFloors+1; i++ {
		if nFloors == 1 {
			Tower.WriteString("*")
		} else {
			for j := i; j <= nFloors-1; j++ {
				Tower.WriteString(" ")
			}
			for k := 1; k < i*2; k++ {
				Tower.WriteString("*")
			}
			for j := i; j <= nFloors-1; j++ {
				Tower.WriteString(" ")
			}
		}

		res = append(res, Tower.String())
		Tower.Reset()

	}
	return res
}
