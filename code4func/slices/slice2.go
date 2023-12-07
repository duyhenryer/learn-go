package main

import "fmt"

func main() {
	// Array
	myArray := [5]int{1, 2, 3, 4, 5}

	// Slice (tạo từ array)
	mySlice := myArray[1:4]

	fmt.Println("Array:", myArray)
	fmt.Println("Slice:", mySlice)

	// Thay đổi giá trị trong slice
	mySlice[0] = 99

	fmt.Println("Array sau khi thay đổi slice:", myArray)
	fmt.Println("Slice sau khi thay đổi:", mySlice)
}
