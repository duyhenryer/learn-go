package main

import "fmt"

/*
Interface
Multiple interface
Embed interface :  Kieu gom nhiều interface lại
Empty interface
*/
type Movement interface {
	move()
}
type Animal interface {
	speak() // method
}

// Embed interface
type NextAnimal interface {
	Movement
	Animal
}

type Dog struct {
}

func (d Dog) speak() {
	fmt.Println("Woow")
}

func (d Dog) move() {
	fmt.Println("Dogggggg")
}

// Empty interface
type data struct {
	index int
}

func goout(i interface{}) {
	fmt.Println(i)
}
func main() {
	var animal Animal
	animal = Dog{}
	animal.speak()
	fmt.Println("------------------")
	dog := Dog{}
	var m Movement = dog
	m.move()
	var a Animal = dog
	a.speak()

	fmt.Println("------------------")
	var na NextAnimal = dog
	na.move()
	na.speak()

	fmt.Println("------------------")
	goout(19)
	goout(10.11)

	d := data{123}
	fmt.Println(d)
}
