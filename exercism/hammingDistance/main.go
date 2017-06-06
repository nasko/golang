package main

import "fmt"

func main() {
	var first, second string

	fmt.Println("Please enter a string:")
	fmt.Scan(&first)

	fmt.Println("Please enter another string:")
	fmt.Scan(&second)

	hamming := func(first, second string) int {
		var distance int
		rFirst := []rune(first)
		rSecond := []rune(second)

		lFirst := len(rFirst)
		lSecond := len(rSecond)
		var rLonger, rShorter []rune

		if lFirst > lSecond {
			rLonger = rFirst
			rShorter = rSecond
		} else {
			rLonger = rSecond
			rShorter = rFirst
		}

		difference := len(rLonger) - len(rShorter)

		for i,_ := range rShorter {
			if string(rShorter[i]) != string(rLonger[i]) {
				distance++
			}
			fmt.Printf("%v - %v\n", string(rShorter[i]), string(rLonger[i]))
		}

		distance += difference
		return distance
	}

	fmt.Println("The hamming distance between the two is ", hamming(first,second))
}

