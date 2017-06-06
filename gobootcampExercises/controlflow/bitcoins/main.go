package main

import "fmt"

var (
	coins = 50
	users = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie",
		"Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
	}
	distribution = make(map[string]int, len(users))
)

func distribute() {
	for _, name := range users {
		sRunes := []rune(name)
		for _, r := range sRunes {
			if distribution[name] == 10 {
				break
			}
			var bonus int
			char := string(r)
			switch char {
			case "a", "e", "A", "E":
				bonus = 1
			case "i", "I":
				bonus = 2
			case "o", "O":
				bonus = 3
			case "u", "U":
				bonus = 4
			default:
				bonus = 0
			}

			if newCoins := distribution[name] + bonus; newCoins > 10 {
				bonus = 10 - distribution[name]
			}

			distribution[name] += bonus
			coins -= bonus
		}
	}
}

func main() {
	distribute()
	fmt.Println(distribution)
	fmt.Println("Coins left:", coins)
}
