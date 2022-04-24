package main

import "fmt"

type Animal struct {
	Name         string
	Sound        string
	NumberOfLegs int
}

func main() {

	// Calling a variadic function
	myTotal := sumMany(2, 3, 4, 5)
	fmt.Println(myTotal)

	// Caling a function attached to the Animal type
	var dog Animal
	dog.Name = "dog"
	dog.Sound = "woof"
	dog.NumberOfLegs = 4
	dog.Says()

	cat := Animal{
		Name:         "cat",
		Sound:        "meow",
		NumberOfLegs: 4,
	}
	cat.Says()
}

// Variadic function
func sumMany(nums ...int) int {
	total := 0

	for _, x := range nums {
		total = total + x
	}

	return total
}

// Function that will be assigned to a type
// The receiver is a pointer to the Animal type
func (a *Animal) Says() {
	fmt.Printf("A %s says %s", a.Name, a.Sound)
}
