package main

import "fmt"

// Input: nums = [3,2,2,3], val = 3
// Output: 2, nums = [2,2,_,_]

func removeElement(nums []int, val int) int {

	p := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[p] = nums[i]
			p++

		}

	}
	return p

}
func main() {
	// Test case.......................
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	val := 3
	k := removeElement(nums, val)
	fmt.Println("k =", k) // Output: k = 7
	fmt.Println("nums =", nums[:k])
}
