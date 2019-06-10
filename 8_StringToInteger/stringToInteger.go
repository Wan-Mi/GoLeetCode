package main

import "math"

func main() {
	println(myAtoi("+0 123"))
	println(myAtoi("   -42"))
	println(myAtoi("9223372036854775808"))
}

func myAtoi(str string) int {

	byteList := []byte{}
	for _, value := range []byte(str) {
		if value != ' ' || len(byteList) > 0 {
			byteList = append(byteList, value)
		}
	}

	if len(byteList) == 0 {
		return 0
	}

	symbol := 1
	if byteList[0] == '-' {
		symbol = -1
	}
	if byteList[0] == '+' || byteList[0] == '-' {
		if len(byteList) == 1 {
			return 0
		}
		byteList = byteList[1:]
	}

	numList := []int{}
	for _, value := range byteList {
		if value < '0' || value > '9' {
			break
		}
		numList = append(numList, int(value-'0'))
	}

	reNum := 0
	maxNum := int(math.Pow(2, 31) - 1)
	for _, value := range numList {
		reNum = reNum*10 + value
		if reNum > maxNum {
			reNum = maxNum
			if symbol < 0 {
				reNum++
			}
			break
		}
	}

	return reNum * symbol
}
