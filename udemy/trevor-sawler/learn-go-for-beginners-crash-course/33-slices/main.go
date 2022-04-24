package main

import (
	"fmt"
	"sort"
)

func main() {
	var animals []string
	animals = append(animals, "dog")
	animals = append(animals, "fish")
	animals = append(animals, "cat")
	animals = append(animals, "horse")

	fmt.Println(animals)
	fmt.Println("")

	printSliceIndexedElements(animals)
	fmt.Println("")

	sort.Strings(animals)

	printSliceIndexedElements(animals)
	fmt.Println("")

	fmt.Println(animals)
	fmt.Println("")

	animals = DeletefromSlice(animals, 1)

	printSliceIndexedElements(animals)
	fmt.Println("")

	fmt.Println(animals)
	fmt.Println("")
}

func DeletefromSlice(a []string, i int) []string {

	fmt.Println("Is it sorted? (1)", sort.StringsAreSorted(a))
	a[i] = a[len(a)-1]

	fmt.Println("Is it sorted? (2)", sort.StringsAreSorted(a))

	a = a[:len(a)-1]

	fmt.Println("Is it sorted? (3)", sort.StringsAreSorted(a))

	return a
}

func printSliceIndexedElements(a []string) {
	for i, x := range a {
		fmt.Printf("animals[%d]: %s\n", i, x)
	}
}
