package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) {
	num1Idx := m - 1
	num2Idx := n - 1
	margedIdx := m + n - 1

	for num1Idx >= 0 && num2Idx >= 0 {
		if nums1[num1Idx] > nums2[num2Idx] {
			nums1[margedIdx] = nums1[num1Idx]
			num1Idx--

		} else {
			nums1[margedIdx] = nums2[num2Idx]
			num2Idx--

		}
		margedIdx--
	}
	for num2Idx >= 0 {
		nums1[margedIdx] = nums2[num2Idx]

		num2Idx--
		margedIdx--

	}

}

func main() {
	//test....

	nums1 := []int{1, 3, 2, 0, 0, 0}
	m := 3
	nums2 := []int{2, 5, 6}
	n := 3
	merge(nums1, m, nums2, n)
	fmt.Println(nums1)

}
