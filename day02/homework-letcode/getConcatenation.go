package main

func getConcatenation(nums []int) []int {
	//return append(nums, nums...)
	result := nums
	for i := 0; i < len(nums); i++ {
		result = append(result, nums[i])
	}
	return result
}

func main() {
	//nums := [...]int{1, 2, 3, 4, 5}
	//var nums [5]int = [...]int{1, 2, 3, 4, 5}
	//fmt.Println(getConcatenation(nums)

}
