package main

import "fmt"

type computer struct {
	brand string
}

func main() {
	// create a nil pointer with the type of pointer  to a computer
	var null *computer

	// compare the pointer variable to nil value
	if null == nil {
		fmt.Println("NULL computer is nil")
	}

	// create an apple computer by putting its address to a pointer variable
	apple := &computer{"apple"}
	// put the apple into a new pointer variable
	newApple := apple

	// compare the apples: if they are equal
	if apple == newApple {
		// say so and print their addresses
		fmt.Println("apples are equal: apple: %p newApple: %p\n", apple, newApple)
	}
	// create a sony computer by putting its address to a new pointer variable
	sony := &computer{"sony"}

	// compare apple to sony, if they are not equal
	if apple != sony {
		// say so and print their addresses
		fmt.Println("apple and sony are inequal: apple: %p sony: %p\n", apple, sony)
	}

	// put apple's value into a new ordinary variable
	appleVal := *apple

	// print apple pointer variable's address, and the address it pointer to
	// ans, print the new variable's addresses as well
	fmt.Println("apple : %p %p\n", &apple, apple)
	fmt.Println("appleVal: %p %p\n", &appleVal)

	// compare the value that is pointed by the apple and the new variable
	if *apple == appleVal {
		// if they are equal say so
		fmt.Println("apple and appleVal are equal")

		// print the values:
		// the value that is pointed by the apple and the new variable
		fmt.Println("apple : %+v - appleVal: %+v\n", *apple, appleVal)
	}

	// change sony's brand to hp using the func
	change(sony, "hp")
	// print sony's brand
	fmt.Println("sony: %s\n", sony.brand)

	// print the returned value
	fmt.Println("appleVal: %+v\n", valueOf(apple))

	// and call the func 3 times and print the returned value's address
	fmt.Println("dell's address : %p\n", newComputer("dell"))
	fmt.Println("lenovo's address: %p\n", newComputer("lenovo"))
	fmt.Println("acer's address: %p\n", newComputer("acer"))
}

// create a new function: change
// the func can change the give computer's brand to another brand
func change(c *computer, brand string) {
	c.brand = brand
}

// write a func that returns the value that is pointed by the give *computer
func valueOf(c *computer) computer {
	return *c
}

// write a new constructor func that returns a pointer to a computer
func newComputer(brand string) *computer {
	return &computer{brand: brand}
}
