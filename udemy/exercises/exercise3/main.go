package main

import "fmt"

// Create a program that prints to the terminal asking for a user to enter their name.
// Use fmt.Scan to read a user’s name entered at the terminal.
// Print “Hello <NAME>” with <NAME> replaced with what the user entered at the terminal.

func main() {
    fmt.Println("Please enter your name:")
    var name string

    fmt.Scan(&name)
    fmt.Println("Hello", name)
}
