package main

import "fmt"

// Step 1: Define an interface named "Speaker"
type Speaker interface {
	Speak() string
}

// Step 2: Implement the Speaker interface for a Dog
type Dog struct {
	Name string
}

// Implement the Speak method for Dog
func (d Dog) Speak() string {
	return fmt.Sprintf("Woof! My name is %s.", d.Name)
}

// Step 3: Implement the Speaker interface for a Cat
type Cat struct {
	Name string
}

// Implement the Speak method for Cat
func (c Cat) Speak() string {
	return fmt.Sprintf("Moew~ My name is %s.", c.Name)
}

// Step 4: Function that takes a Speaker interface and makes it speak
func makeSpeak(speaker Speaker) {
	fmt.Println(speaker.Speak())
}

func main() {
	// Step 5: Create instances of Dog and Cat
	dog := Dog{Name: "Buddy"}
	cat := Cat{Name: "Kities"}
	fmt.Println(dog)
	fmt.Println(cat)

	name := "John"
	age := 30

	// Sprintf formats the string and returns it
	formattedString := fmt.Sprintf("Name: %s, Age: %d", name, age)
	fmt.Println(formattedString)
}
