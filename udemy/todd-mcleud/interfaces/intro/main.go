package main

import "fmt"

type Square struct {
	side int
}

func (s Square) area() int {
	return s.side * s.side
}

type Shape interface {
	area() int
}

func info(s Shape) {
	fmt.Println(s)
	fmt.Println(s.area())
}

func main() {
	s := Square{10}
	info(s)
}
