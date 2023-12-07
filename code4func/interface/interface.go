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

type Dog1 struct {
}

func (d Dog1) speak() {
	fmt.Println("Woow")
}

func (d Dog1) move() {
	fmt.Println("Dog1ggggg")
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
	animal = Dog1{}
	animal.speak()
	fmt.Println("------------------")
	Dog1 := Dog1{}
	var m Movement = Dog1
	m.move()
	var a Animal = Dog1
	a.speak()

	fmt.Println("------------------")
	var na NextAnimal = Dog1
	na.move()
	na.speak()

	fmt.Println("------------------")
	goout(19)
	goout(10.11)

	d := data{123}
	fmt.Println(d)
}
