// LeetCode Problem 155: Min Stack
//
// Design a stack that supports push, pop, top, and retrieving the minimum element in constant time.
//
// Implement the MinStack class:
//
// - MinStack() initializes the stack object.
// - void push(int val) pushes the element val onto the stack.
// - void pop() removes the element on the top of the stack.
// - int top() gets the top element of the stack.
// - int getMin() retrieves the minimum element in the stack.
//
// You must implement a solution with O(1) time complexity for each function.
//
// Example 1:
// Input
// ["MinStack","push","push","push","getMin","pop","top","getMin"]
// [[],[-2],[0],[-3],[],[],[],[]]
//
// Output
// [null,null,null,null,-3,null,0,-2]
//
// Explanation
// MinStack minStack = new MinStack();
// minStack.push(-2);
// minStack.push(0);
// minStack.push(-3);
// minStack.getMin(); // return -3
// minStack.pop();
// minStack.top();    // return 0
// minStack.getMin(); // return -2
//
// Constraints:
// - -2^31 <= val <= 2^31 - 1
// - Methods pop, top and getMin operations will always be called on non-empty stacks.
// - At most 3 * 10^4 calls will be made to push, pop, top, and getMin.
//
// Approach Explanation:
// - We use two stacks: one main stack for all values, and an auxiliary minStack to track the current minimum at each level.
// - Push: Add to main stack. If value <= current min (or minStack empty), add to minStack.
// - Pop: Remove from main. If removed value == current min, remove from minStack.
// - Top: Peek at main stack top.
// - GetMin: Peek at minStack top (always the current min).
// - Why <=? To handle duplicates—pushing the same min multiple times requires tracking each instance, so popping doesn't prematurely remove the min.
// - Time: O(1) for all operations.
// - Space: O(n) in worst case (e.g., decreasing order).

package main

import "fmt"

type MinStack struct {
	main     []int
	minstack []int
}

func Constructor() MinStack {
	return MinStack{
		main:     []int{},
		minstack: []int{},
	}

}

func (this *MinStack) Push(val int) {
	this.main = append(this.main, val)

	if len(this.minstack) == 0 || val <= this.minstack[len(this.minstack)-1] {
		this.minstack = append(this.minstack, val)
	}
}
func (this *MinStack) Pop() {
	if len(this.main) == 0 {
		return // Assuming valid calls as per constraints, but handle empty for safety
	}

	top := this.main[len(this.main)-1]
	this.main = this.main[:len(this.main)-1]

	if top == this.minstack[len(this.minstack)-1] {
		this.minstack = this.minstack[:len(this.minstack)-1]

	}
}
func (this *MinStack) Top() int {
	if len(this.main) == 0 {
		return 0 // Constraints say non-empty, but return 0 or handle as needed
	}

	return this.main[len(this.main)-1]

}

func (this *MinStack) GetMin() int {
	if len(this.minstack) == 0 {
		return 0 // Constraints say non-empty
	}
	return this.minstack[len(this.minstack)-1]
}

func main() {
	// test
	obj := Constructor()
	obj.Push(-2)
	obj.Push(0)
	obj.Push(-3)

	// Check min
	fmt.Println("Current Min:", obj.GetMin()) //  -3

	// Pop the top
	obj.Pop()

	// Check top and min
	fmt.Println("Current Top:", obj.Top())    // Should be 0
	fmt.Println("Current Min:", obj.GetMin()) // Should be -2

	// Add another push and test
	obj.Push(-1)
	fmt.Println("After pushing -1, Min:", obj.GetMin()) // Should be -2 (wait, no: -1 > -2? Wait, -1 is greater than -2? No: -2 is smaller than -1.
	// Correction: -2 < -1, so min stays -2
	// But let's push a smaller one
	obj.Push(-4)
	fmt.Println("After pushing -4, Min:", obj.GetMin()) // Should be -4

	// Pop and check
	obj.Pop()
	fmt.Println("After pop, Min:", obj.GetMin()) // Back to -2

}
