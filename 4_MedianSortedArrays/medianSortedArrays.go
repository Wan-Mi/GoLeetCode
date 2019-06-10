package main

import "fmt"

func main() {

	fmt.Println("start")
	nums1 := []int{
		100001,
	}

	nums2 := []int{
		100000,
	}

	fmt.Println(findMedianSortedArrays(nums1, nums2))
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	A := nums1 //short
	B := nums2 //long
	if len(nums1) > len(nums2) {
		B = nums1
		A = nums2
	}
	m, n := len(A), len(B)

	for i := m / 2; ; {
		j := (m+n+1)/2 - i

		if i > 0 && A[i-1] > B[j] {
			i--

		} else if i < m && B[j-1] > A[i] {
			i++

		} else {
			fmt.Println(i, j)
			var left, right int

			if i == 0 {
				left = B[j-1]
			} else if j == 0 {
				left = A[i-1]
			} else {
				left = B[j-1]
				if A[i-1] > B[j-1] {
					left = A[i-1]
				}
			}

			if (m+n)%2 > 0 {
				return float64(left)
			}

			if i == m {
				right = B[j]
			} else if j == n {
				right = A[i]
			} else {
				right = B[j]
				if A[i] < B[j] {
					right = A[i]
				}
			}

			return (float64(left) + float64(right)) / 2
		}
	}
}
