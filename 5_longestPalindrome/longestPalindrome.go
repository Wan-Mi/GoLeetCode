package main

func main() {
	println(longestPalindrome("cbbd"))
}

func longestPalindrome(s string) string {
	byteString := []byte(s)

	longestPalindrome := ""
	for subStringEnd, _ := range byteString {
		for subStringStart, _ := range byteString[:subStringEnd+1] {
			if len(byteString[subStringStart:subStringEnd+1]) > len(longestPalindrome) {
				if isPalindrome(byteString[subStringStart : subStringEnd+1]) {
					longestPalindrome = string(byteString[subStringStart : subStringEnd+1])
				}
			}
		}
	}

	return longestPalindrome
}

func isPalindrome(byteString []byte) bool {
	length := len(byteString)
	if length == 1 || (length == 2 && byteString[0] == byteString[1]) {
		return true
	}

	if byteString[0] == byteString[length-1] {
		return isPalindrome(byteString[1 : length-1])
	}

	return false
}
