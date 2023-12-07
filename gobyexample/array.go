package main

import "fmt"

func main() {
	var a [5]int            // Khai báo var + arrayName + Độ dài[tên kiểu]
	fmt.Println(a, "Empty") // Mặt định là 0

	a[4] = 100
	fmt.Println("Set:", a)
	fmt.Println("Get: ", a[4])
	fmt.Println("Len: ", len(a))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println(b)

	names := [3]string{"Duy", "Henry", "Tom"}
	// Using a for loop to iterate
	for i := 0; i <= len(names); i++ {
		fmt.Println("Element at index", i, ":", names[i])
	}

	// Using range in a for loop =>  without worrying about indices.
	for index, value := range names {
		fmt.Println("Element at index", index, ":", value)
	}
	// 2D array
	var twoD [2][3]int
	z := 0
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = z
			//z += 1
			z++
		}
	}
	fmt.Println(twoD)
}
