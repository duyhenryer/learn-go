package main

import "fmt"

func main() {
	// Khai bao slice
	var slice []int
	fmt.Println(slice)

	// Khai bao va khoi tao slice
	var slice1 = []int{1, 2, 3, 4}
	fmt.Println(slice1)

	// Tao 1 slice tu 1 array
	var array = [4]int{1, 2, 3, 4} // Index: 0 1 2 3
	slice2 := array[1:3]           // array[1(với  1 là index) -> array[3-1=2] <=> array[2]
	// index 1: 2
	// index 3 - 1 = 2 => 3
	fmt.Println(slice2)

	slice3 := array[:]
	fmt.Println(slice3)

	slice4 := array[2:] // Lấy từ index =2 trở về sau.
	fmt.Println(slice4)

	slice5 := array[3:] // Lấy từ index =3 trở về sau.
	fmt.Println(slice5)

	// Tao 1 slice từ 1 slice khác
	var slice6 = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	slice7 := slice6
	fmt.Println(slice7)

	slice8 := slice6[1:] // Index =1
	fmt.Println(slice8)

	// Slice => Reference type
	var array1 = [5]int{1, 2, 3, 4, 5} // Khởi tạo 1 size 5 giá trị
	slice9 := array1[:]
	slice9[0] = 777 // Sau khi gán slice9 = 777 thì array nó cũng thay đổi => Slice nó là 1 refer type.
	// Không như Array là value type
	fmt.Println(slice9)
	fmt.Println(array1)
	/*
		[777 2 3 4 5]
		[777 2 3 4 5]
	*/

	// Length va capacity cua slice
	countries := [...]string{"VN", "USA", "CANADA", "FRANCE", "CHINA"}
	slice10 := countries[1:3] // index 1, value: 3-1= 2
	fmt.Println(slice10)

	fmt.Println(len(slice10))
	fmt.Println(cap(slice10)) // Để tính cap, thì lấy tất cả phía sau index 1
	/*
		len: so lượng phần tử của slice
		cap: số lượng phần tử của underlying array bắt đầu từ vị trí start
		khi mà slice được tạo.
		Cái Pointer của slice, nó sẽ trỏ tới underlying array = mảng của countries trên.
		Để tính cap, thì lấy tất cả phía sau index 1 => cap = 4(USA, CANADA, FRANCE, CHINA)
		EX thêm: slice10 := countries[3:4] => index 3 => FRANCE
		=> len = 1 => cap(lấy từ vị trí start của index) = 2
	*/

	fmt.Println("_________________")
	/*
		Các hàm căn bản để làm việc với slice
		Make, copy, append
		- make: Khai báo 1 slice có cụ thể len và cap là bao nhiêu.
		- append:
		- copy:
	*/
	// 1. make
	slice11 := make([]int, 2, 5) // Hàm make này, mình sẽ truyền kiểu dữ liệu cho underlying array
	fmt.Println(slice11)
	fmt.Println(len(slice11))
	fmt.Println(cap(slice11))
	/*
			[0 0] => do minhg chưa khởi tạo gì cho mảng trên.
			2 = len
			5 = cap
		Nếu không khai báo cap cho clice11 thì cap sẽ = len = 2
	*/
	// 2. append
	var slice13 []int // array rỗng
	slice13 = append(slice13, 111111)
	fmt.Println(slice13)

	// 3. copy
	src := []string{"A", "B", "C", "D"}
	dest := make([]string, 2)

	copy(dest, src)

	number := copy(dest, src)
	fmt.Println(number)
	fmt.Println(dest)
	/*
		Return số phần tử dc copy
	*/

	/*
			Trong API của GO thì không có delete trong slice.
			Mình dùng trick của hàm append để delete
		 	Delete item trong 1 slice với index 1 (B)
	*/
	src = append(src[:1], src[2:]...) // Append là nối 2 slice với slice = append(slice1, slice2...)
	fmt.Println(src)

}
