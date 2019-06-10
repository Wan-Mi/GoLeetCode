package main

func main() {
	println(convert("asdfghjk", 3))
	println(convert2("asdfghjk", 3))
}

func convert(s string, numRows int) string {

	if numRows == 1 {
		return s
	}

	reByte := []byte{}
	strMap := map[int]byte{}
	flagS := map[int]byte{}
	virI := int(0)
	for i := 0; i < len([]byte(s)); {
		flagS[i] = []byte(s)[i]
		reByte = append(reByte, []byte(s)[i])
		i += (numRows - 1) * 2
		virI = i
	}
	flagS[virI] = ' '

	for idx, val := range []byte(s) {
		strMap[idx] = val
	}

	for i := 1; i < numRows-1; i++ {
		var key int
		for key = 0; key < len([]byte(s)); key += (numRows - 1) * 2 {
			if byteVal, ok := strMap[key-i]; ok {
				reByte = append(reByte, byteVal)
			}
			if byteVal, ok := strMap[key+i]; ok {
				reByte = append(reByte, byteVal)
			}
		}
		if byteVal, ok := strMap[key-i]; ok {
			reByte = append(reByte, byteVal)
		}
	}

	for i := 0; i < len([]byte(s)); i += (numRows - 1) * 2 {
		if byteVal, ok := strMap[i+numRows-1]; ok {
			reByte = append(reByte, byteVal)
		}
	}

	return string(reByte)
}

// better
func convert2(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	lineString := make([]string, numRows)
	step := 1
	lineNum := 0
	for idx, val := range []byte(s) {

		lineString[lineNum] += string(val)

		if idx%(numRows-1) == 0 && idx != 0 {
			step = step * (-1)
		}
		lineNum += step
	}

	reString := ""
	for _, subString := range lineString {
		reString += subString
	}

	return reString
}
