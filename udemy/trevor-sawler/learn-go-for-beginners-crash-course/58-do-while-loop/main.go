package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/eiannone/keyboard"
)

func main() {

	coffees := make(map[int]string)
	coffees[1] = "Cappucino"
	coffees[2] = "Latte"
	coffees[3] = "Americano"
	coffees[4] = "Mocha"
	coffees[5] = "Macchiato"
	coffees[6] = "Espresso"

	char := ' '
	// do..while variant 1:
	for {
		char, _, err := keyboard.GetSingleKey()
		if err != nil {
			log.Fatal(err)
		}

		if char == 'q' || char == 'Q' {
			break
		}

		i, _ := strconv.Atoi(string(char))

		_, ok := coffees[i]
		if ok {
			fmt.Println(fmt.Sprintf("You chose %s", coffees[i]))
		}
	}

	// do..while variant 2:
	for ok := true; ok; ok = char != 'q' && char != 'Q' {
		char, _, err := keyboard.GetSingleKey()
		if err != nil {
			log.Fatal(err)
		}

		i, _ := strconv.Atoi(string(char))

		_, ok := coffees[i]
		if ok {
			fmt.Println(fmt.Sprintf("You chose %s", coffees[i]))
		}
	}
}
