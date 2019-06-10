package main

import "fmt"

func main() {
	lay := 30
	fmt.Println(getMinTimes(lay))
}

func getMinTimes(layer int) int {
	if layer <= 0 {
		return 0
	}

	if layer == 1 {
		return 1
	}

	min := layer // 一栋layer层的大楼试验次数肯定不可能超过layer次。
	var temp int
	for i := 1; i <= layer; i++ {
		temp = 1 + getMax(i-1, getMinTimes(layer-i))
		if min > temp {
			min = temp

		}
	}

	return min
}

func getMax(i, j int) int {
	if i > j {
		return i
	}
	return j
}
