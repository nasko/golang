package main

import "fmt"

// Create a program that prints to the terminal asking for a user to enter a small number and a larger number.
// Print the remainder of the larger number divided by the smaller number.

func main() {
    var small, big int

    //fmt.Println("Please enter a small number:")
    //fmt.Scan(&small)
    //
    //fmt.Println("Please enter a big number:")
    //fmt.Scan(&big)

    fmt.Println("Please enter a big and a small number, space-delimitted:")
    fmt.Scan(&big, &small)

    //fmt.Printf("The remainder of %v/%v is %v", big, small, big % small)
    fmt.Printf("Big: %v - %T\n", big, big)
    fmt.Printf("Small: %v - %T\n", small, small)

}
