package main

import "fmt"

func main() {
	//Declare
	var b *int
	a :=2
	b = &a

	fmt.Println(b)
	fmt.Println(*b)

	b = new(int)
    *b = 10
    fmt.Println(*b)
    fmt.Println(b)
}