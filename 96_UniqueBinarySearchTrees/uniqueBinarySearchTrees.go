package main

import "fmt"

func main() {
	println(numTrees(3))
}

func numTrees(n int) int {

	dpNum := make([]int, n+1)
	dpNum[0] = 1
	dpNum[1] = 1
	for i := 2; i <= n; i++ {
		for j := 1; j <= i; j++ {
			dpNum[i] += dpNum[j-1] * dpNum[i-j]
		}
	}

	fmt.Println(dpNum)
	return dpNum[n]
}
