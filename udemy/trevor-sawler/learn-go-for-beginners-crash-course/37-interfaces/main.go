package main

import (
	"fmt"
)

// interface type
type Animal interface {
	Says() string
	HowManyLegs() int
}

// Some custom types
type Dog struct {
	Name     string
	Sound    string
	NrOfLegs int
}

// Now, after we attach these two functions to the Dog type, it will satisfy the Animal interface
func (d *Dog) Says() string {
	return d.Sound
}
func (d *Dog) HowManyLegs() int {
	return d.NrOfLegs
}

type Cat struct {
	Name     string
	Sound    string
	NrOfLegs int
	HasTail  bool
}

func (c *Cat) Says() string {
	return c.Sound
}
func (c *Cat) HowManyLegs() int {
	return c.NrOfLegs
}

func main() {
	// ask a riddle
	dog := Dog{
		Name:     "dog",
		Sound:    "woof",
		NrOfLegs: 4,
	}

	// We need to pass a refference to dog, because the two functions
	// attached to the Dog type have a pointer receiver (*Dog)
	Riddle(&dog)

	cat := Cat{
		Name:     "cat",
		Sound:    "meow",
		NrOfLegs: 4,
		HasTail:  true,
	}

	Riddle(&cat)
}

func Riddle(a Animal) {
	riddle := fmt.Sprintf(`This anymal says "%s" and has %d legs. What animal is it?`, a.Says(), a.HowManyLegs())
	fmt.Println(riddle)
}



