package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(FindShort("this is a simple test"))                        // 1
	fmt.Println(FindShort("cat bat rat hat"))                              // 3
	fmt.Println(FindShort("find  the    shortest    word"))                // 3
	fmt.Println(FindShort("this is a supercalifragilisticexpialidocious")) // 1

}

func FindShort(s string) int {
	words := strings.Fields(s)
	shortestWord := len(words[0])

	for _, v := range words {
		if len(v) < shortestWord {
			shortestWord = len(v)
		}
	}

	return shortestWord

}
