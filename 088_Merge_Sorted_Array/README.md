```markdown
# 88. Merge Sorted Array

## Problem

[LeetCode Problem Link](https://leetcode.com/problems/merge-sorted-array/)

You are given two integer arrays `nums1` and `nums2`, sorted in non-decreasing order, and two integers `m` and `n`, representing the number of elements in `nums1` and `nums2` respectively.

Merge `nums1` and `nums2` into a single array sorted in non-decreasing order. The final sorted array should be stored in `nums1`, which has a length of `m + n`, where the first `m` elements are the elements to be merged, and the last `n` elements are set to 0 and should be ignored.

### Example

**Input:** `nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3`  
**Output:** `[1,2,2,3,5,6]`

## Solution

- **Language**: Go
- **File**: [solution.go](solution.go)

### Approach

- Use three pointers:
  - `num1Idx` for the last valid element in `nums1` (index `m-1`).
  - `num2Idx` for the last element in `nums2` (index `n-1`).
  - `margedIdx` for the current position in `nums1` (index `m+n-1`).
- Merge from the end of both arrays to avoid overwriting elements in `nums1`.
- Compare `nums1[num1Idx]` and `nums2[num2Idx]`, placing the larger element at `nums1[margedIdx]` and decrementing the appropriate pointers.
- If `nums2` has remaining elements, copy them to `nums1`.

### Time and Space Complexity

- **Time Complexity**: O(m + n), where `m` and `n` are the lengths of `nums1` and `nums2`.
- **Space Complexity**: O(1), as the merge is done in-place.
```
