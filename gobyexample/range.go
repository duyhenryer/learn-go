package main

import "fmt"

func main() {
	nums := [4]int{1, 2, 3, 4}

	for index, value := range nums {
		fmt.Println("Index %d , value %d: ", index, value)
	}
}
