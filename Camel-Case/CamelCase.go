package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "hi there"
	str2 := "hey there, how are you doing?"
	function := CamelCase(str)
	function2 := CamelCase(str2)

	fmt.Println(function)
	fmt.Println(function2)
}

func CamelCase(s string) string {
	words := strings.Split(s, "  ")
	key := ""
	for _, word := range words[0:] {
		key += strings.Title(word)
	}
	return key
}
