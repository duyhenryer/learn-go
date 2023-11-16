package main

import "fmt"

func countKDifference(nums []int, k int) int {
	count := 0
	// for i := 0; i < len(nums); i++
	for i := range nums { // i = index, e = element
		for j := 0; j < i; j++ {
			if nums[i]-nums[j] == k || nums[i]-nums[j] == -k {
				count++
			}

		}
	}
	return count
}

func main() {
	nums := []int{1, 2, 2, 1}
	k := 1
	countKDifference(nums, k)
	fmt.Println(countKDifference(nums, k))
}
