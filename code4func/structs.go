package main

import "fmt"

type StudentInfo struct {
	class string
	email string
	age   int
}

type Student struct {
	id   int
	name string
	info StudentInfo
}

func main() {
	// named
	st1 := Student{
		id:   123,
		name: "Ryan",
	}
	fmt.Println(st1)
	fmt.Println(st1.id)
	fmt.Println(st1.name)
	/*
		st2 := Student{456, "Robin"}
		fmt.Println(st2)
		fmt.Println(st2.id)
		fmt.Println(st2.name)

		var st3 Student = struct {
			id   int
			name string
		}{id: 777, name: "DuyHenry"}
		fmt.Println(st3)
	*/

	// Ở trên có 3 cách khai báo cơ bản.

	var anonymous = struct {
		email string
		age   int
	}{email: "duyne@duy.com",
		age: 25}
	fmt.Println(anonymous)

	// Pointer trỏ tới struct
	pointer := &Student{id: 999,
		name: "Robin"}
	fmt.Println(&pointer)
	fmt.Println((*pointer).id) // Có dấu * là khai báo tường minh.
	fmt.Println(pointer.id)    // Không có * là ngắn gọn.
	fmt.Println((*pointer).name)
	fmt.Println(pointer.name)

	// anonymous fields: Không cần tên, chỉ cần kiểu dữ liệu: string
	type NoName struct {
		string
		int
	}

	n := NoName{"DuyHenry", 9999}
	fmt.Println(n)

	// Struct lồng struct
	st2 := Student{456, "Robin", StudentInfo{class: "cs-01", email: "duyyy@mail.com", age: 111}}
	fmt.Println(st2)

	// Compare 2 Struct
	type struct1 struct {
		id   int
		name string
		//info map[int]int // info có kiểu dữ liệu là map, key là int, kiểu dữ liệu là int
	}
	//s1 := struct1{1, "A", map[int]int{0: 1}}  ***
	//s2 := struct1{1, "B ", map[int]int{0: 1}}
	s1 := struct1{1, "A"}
	s2 := struct1{1, "A"}

	if s1 == s2 {
		fmt.Println("s1 == s2")
	} else {
		fmt.Println("s1 != s2")
	}
	// *** => Không thể so sánh đc, khi có 1 kiểu dữ liệu Map nó không thể so sánh => không ss dc,
	// Nếu bỏ kiểu dữ liệu Map thì ss được

	// Zero-value => chừng nào chưa khởi tạo thì nó lấy = 0
	// Trong đó mà có lồng struct thì cũng = 0
	var student Student
	fmt.Println(student)

}
