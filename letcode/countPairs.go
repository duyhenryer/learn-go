// https://leetcode.com/problems/count-pairs-whose-sum-is-less-than-target/description/
package main

import (
	"fmt"
	"sort"
)

func countPairs(nums []int, target int) int {
	sort.Ints(nums) //sort the vector nums
	count := 0
	left := 0
	right := len(nums) - 1

	for left < right { // loop until left is less than righ
		if nums[left]+nums[right] < target {
			count += right - left
			left++ // increment the left
		} else {
			right-- // decrement the right
		}
	}
	return count

}
func main() {
	nums1 := []int{-1, 1, 2, 3}
	target1 := 2
	out := countPairs(nums1, target1)
	fmt.Println(out)
}
