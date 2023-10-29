package main

import (
	"strconv"
)

func SumMix(arr []any) int {
	sum := 0

	for _, v := range arr {
		switch v.(type) {
		case int:
			sum += v.(int)
		case string:
			s, _ := strconv.Atoi(v.(string))
			sum += s
		}
	}
	return sum
}
