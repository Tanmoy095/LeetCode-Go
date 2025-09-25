// threeSum finds all unique triplets in the array which sum to zero.
// Approach:
// 1. Sort the input array to make it easier to avoid duplicates and use two pointers.
// 2. Iterate through the array, treating each element as a potential first element of a triplet.
// 3. For each first element, use two pointers (left and right) to find pairs that sum with the first element to zero.
// 4. Skip duplicate elements to ensure unique triplets.
// 5. Move pointers inward, skipping duplicates, and collect valid triplets.
// Time Complexity: O(n^2), Space Complexity: O(n) for the result.
package main

import (
	"fmt"
	"slices"
)

// threeSum finds all unique triplets in nums that sum to zero.
func threeSum(nums []int) [][]int {
	slices.Sort(nums)
	result := [][]int{}

	for i := 0; i < len(nums)-2; i++ {
		// Skip duplicate values for the first element
		if i > 0 && nums[i] == nums[i-1] {
			continue //if duplicate found - Skip everything below and go to nextr i
		}
		left := i + 1
		right := len(nums) - 1
		// Use two pointers to find pairs that sum to zero with nums[i]
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum < 0 {
				left++ // Need a larger sum
			} else if sum > 0 {
				right-- // Need a smaller sum
			} else {
				// Found a valid triplet
				result = append(result, []int{nums[i], nums[left], nums[right]})
				// Skip duplicates for left pointer
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				// Skip duplicates for right pointer
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			}
		}
	}
	return result
}

func main() {

	numbers := []int{-4, -1, -1, 0, 1, 2}

	result := threeSum(numbers)
	fmt.Println("Three sum results:", result)
}
