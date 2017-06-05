package main

import (
	"fmt"
)

func filter(in []int, callable func(i int) int) []int {

	xs := []int{}

	for _,n := range in {
		xs = append(xs, callable(n))
	}
	return xs
}

func main() {
	out := filter([]int{1,2,3,4,5,6,7,8,9,10}, func(n int) int{
		return n * 2
	})

	fmt.Println(out)
}
