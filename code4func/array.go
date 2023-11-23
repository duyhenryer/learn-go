package main

import "fmt"

func main() {
	// 1. Khai bao array
	var myArray [4]int // 4 phan tu
	fmt.Println(myArray)
	// 2. Gan gia tri cho array
	myArray[0] = 123
	fmt.Println(myArray)

	myArray[3] = 456
	fmt.Println(myArray)

	// 3. Vua khai bao, vua khoi tao gia tri
	arrays1 := [3]int{1, 2, 3}
	fmt.Println(arrays1)

	arrays2 := [3]int{100}
	fmt.Println(arrays2)

	// 4. Size mang
	fmt.Println(len(arrays2))

	// 5. Khai bao mang khong can set size
	arrays3 := [...]string{"Vinfast", "Honda", "KIA", "Tesla"}
	fmt.Println(arrays3)
	fmt.Println(len(arrays3))
	fmt.Println(arrays3[2]) // Lay index = 2 => KIA

	// 6. Array la value type, khong phai la reference type
	countries := [...]string{"VN", "CHINA", "USA"}
	copyCountries := countries // Gán vùng nhớ cho copyCountries

	copyCountries[0] = "CANADA"
	// sau khi thêm CANADA, thì copyCountries thay đổi trong mảng,
	// Còn countries thì không thay đổi gì
	fmt.Println(countries)
	fmt.Println(copyCountries)

	// 7. Loop array
	for i := 0; i < len(countries); i++ {
		fmt.Println(countries[i])
	}
	for index, value := range countries {
		fmt.Println(index, value)
	}
	// Chỉ muốn làm việc với value, bỏ qua index và ngược lại
	for _, value := range countries {
		fmt.Println(value)
	}
	for index, _ := range countries {
		fmt.Println(index)
	}

}
