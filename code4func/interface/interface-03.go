package main

import "fmt"

// Interface Embedding

// Walker interface
type Walker interface {
	Walk() string
}

// Swimmer interface
type Swimmer interface {
	Swim() string
}

// Human interface embedding Walker and Swimmer interfaces
type Human interface {
	Walker
	Swimmer
	Talk() string // Additional method in the Human interface
	GetAge() int  // New method related to age
}

// Person struct implementing the Human interface
type Person struct {
	Name string
	Age  int
}

// Walk method for Person
func (p Person) Walk() string {
	return fmt.Sprintf("%s is walking", p.Name)
}

// Swim method for Person
func (p Person) Swim() string {
	return fmt.Sprintf("%s is swimming", p.Name)
}

// Talk method for Person
func (p Person) Talk() string {
	return fmt.Sprintf("%s say hello to me", p.Name)
}

func (p Person) GetAge() int {
	return p.Age
}

func main() {
	// Create a Person instance
	person := Person{Name: "Duy Henry", Age: 26}
	// Call methods from the interfaces
	fmt.Println(person.Walk())
	fmt.Println(person.Swim())
	fmt.Println(person.Talk())
	// Access the new method related to age
	fmt.Println("Age: ", person.GetAge())
}
