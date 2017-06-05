package main

import "fmt"

type person struct {
	first string
	last string
	age int
}

func (p person) fullName() string {
	return p.first + " " + p.last
}

func main() {
	nasko := person{"Atanas", "Vasilev", 44}
	eva := person{"Eva", "Vasileva", 10}

	fmt.Println(nasko.fullName())
	fmt.Println(eva.fullName())
}
