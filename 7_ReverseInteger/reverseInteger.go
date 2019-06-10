package main

import "math"

func main() {
	println(reverse(1234567))
	println(reverse(1534236469))
}

func reverse(x int) int {
	symbol := 1
	if x < 0 {
		symbol = -1
		x = -x
	}

	numList := []int{}
	for ; x > 0; x = x / 10 {
		numList = append(numList, x%10)
	}

	reX := 0
	for _, val := range numList {
		reX = reX*10 + val
	}

	if reX > int(math.Pow(2, 31)-1) {
		return 0
	}

	return reX * symbol
}
