package main

import "fmt"

func main() {
	i, j := 42, 2701

	p := &i         // Gía trị của P là địa chỉ của i
	fmt.Println(p)  // Ra hash 0x140000a6018. 0xx là giá trị ô nhớ.
	fmt.Println(*p) // in giá trị ô nhớ của biến p đang trỏ tới i
	*p = 21         // Địa chỉ P có ô nhớ 0x1400010e008 trên RAM. Ta gán giá trị địa chỉ P = 21
	fmt.Println(i)

	p = &j       // Lần này p trỏ tới biến j, không còn trỏ tới biến j
	*p = *p / 37 // *p bên phải, là nó lấy giá trị ô nhớ biến p đang trở ở trên line 10
	// nhưng nó đã được trỏ tới j là dòng 14.
	// => *p = 2701 / 37
	fmt.Println(p)
	fmt.Println(*p) // Lúc này ô nhớ p đã thay đổi.
	/*
		&: Là Address(địa chỉ)
		*: Giá trị
	*/
	t := T{}
	f(&t)
	fmt.Println(t.X)
}

type T struct {
	X int
	Y int
}

func f(t *T) int {
	t.X = 1
	return t.X + t.Y
}
