package main

func main() {
	println(isPalindrome(12321))
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	numList := []int{}
	for ; x > 0; x = x / 10 {
		numList = append(numList, x%10)
	}

	for len(numList) > 1 {
		if numList[0] != numList[len(numList)-1] {
			return false
		}
		numList = numList[1 : len(numList)-1]
	}

	return true
}
