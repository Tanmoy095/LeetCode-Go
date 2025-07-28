package main

import "fmt"

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0

	}

	k := 1

	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[k-1] {
			nums[k] = nums[i]
			k++

		}
	}
	return k

}
func main() {
	nums := []int{1, 1, 2, 2, 3, 3, 4}
	k := removeDuplicates(nums)
	fmt.Println("k=", k)
	fmt.Println("nums =", nums[:k])

}
