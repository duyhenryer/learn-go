package main

import "fmt"

func findMaxK(nums []int) int {
	s := make(map[int]bool)
	k := -1

	// Duyệt qua mảng số nguyên đầu vào và điền giá trị vào map
	// Nếu bỏ _ thì v sẽ hiểu là index.
	// Không cần dùng biến gì. kiểu return 3 số, mà dùng 1 số thì để _
	for _, v := range nums {
		s[v] = true
		//s[nums[i]] = true
		// fmt.Println("Value of s:", s)
	}
	for v := range s {
		if s[-v] && k < v {
			k = v
		}
	}
	return k

}

func main() {
	nums := []int{-1, 10, 6, 7, -7, 1}
	findMaxK(nums)
	fmt.Println(findMaxK(nums))
}
