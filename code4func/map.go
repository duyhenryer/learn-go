package main

import "fmt"

func main() {
	// Key value
	var myMap = make(map[string]int) // string" kieu du lieu: key, int: kieu du lieu value
	fmt.Println(myMap)

	if myMap == nil {
		fmt.Println("myMap = nil")
	} else {
		fmt.Println("myMap =! nil")
	}
	// myMap này không còn nil nữa. Vì nó đã khỏi tạo trong bộ nhớ
	fmt.Println("------------")
	var myMap1 map[string]int
	fmt.Println(myMap1)

	if myMap1 == nil {
		fmt.Println("myMap1 = nil")
	}
	// myMap1: Khởi tạo ban đầu là nil

	// Khai báo với giá trị khởi tạo.
	myMap2 := map[string]int{
		"key1": 1,
		"key2": 2,
		"key3": 3,
	}
	fmt.Println(myMap2)

	// thêm phần tử vào Map
	myMap2["key4"] = 4
	myMap2["key5"] = 5
	fmt.Println(myMap2)

	// delete 1 phaanf tuwr trong map = delete(map, key)
	delete(myMap2, "key1")
	fmt.Println(myMap2)

	// Len của map
	fmt.Println(len(myMap2))

	// Map là 1 reference type
	myMap3 := myMap2
	fmt.Println(myMap3)
	myMap3["key5"] = 100000
	fmt.Println(myMap2)
	// Key5 ban đầu là 5, sau khi myMap3 thay đổi thì ở myMap2 cũng thay đổi thành 10000

	// Truy cập 1 phần tử trong map
	value, found := myMap2["key2"]
	// để xử lý value, thì trước tiên cần kiểm trên
	// tra trong myMap2 có key: key2, thì thêm found để xử lí trước.
	if found {
		fmt.Println(value)
	} else {
		fmt.Println("Khong tim thay gia tri voi key = 10000")
	}

	// Trong Map không có toán tử == => cần viết func để kiểm tra
	//if myMap2 == myMap3 {
	//	fmt.Println(value)
	//} => NO
}
