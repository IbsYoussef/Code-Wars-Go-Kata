package main

import (
	"strconv"
)

func StringToNumber(str string) int {
	intNum, _ := strconv.Atoi(str)
	return intNum
}
