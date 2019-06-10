package main

/**
	数组数字相当于模板高度，间隔1 求任意两根木板间最大储水量
 */
func main() {
	height := []int{1, 2, 3, 4, 5}
	println(maxArea(height))
	println(maxArea1(height))

}

/**
	遍历每根木板
 */
func maxArea(height []int) int {

	maxArea := 0
	for i := 0; i < len(height); i++ {
		for j := i + 1; j < len(height); j++ {
			low := height[i]
			if height[j] < low {
				low = height[j]
			}

			if (j-i)*low > maxArea {
				maxArea = (j - i) * low
			}
		}
	}

	return maxArea
}


/**
	双指针，从两端往中间遍历
 */
//better
func maxArea1(height []int) int {

	maxArea := 0
	for i, j := 0, len(height)-1; i < j; {
		low := height[i]
		if height[j] < low {
			low = height[j]
		}

		if (j-i)*low > maxArea {
			maxArea = (j - i) * low
		}

		if low == height[i] {
			i++
		} else {
			j--
		}
	}

	return maxArea
}
