package main

import "fmt"

func hello() {
    fmt.Println("Hello")
}
func world() {
    fmt.Println("world")
}
func one() {
    fmt.Println(1)
}
func two() {
    fmt.Println(2)
}
func three() {
    fmt.Println(3)
}

func main() {
    defer world()
    defer one()
    defer two()
    hello()

}
