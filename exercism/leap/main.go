package main

import "fmt"

func main() {
	var number int
	for {
		fmt.Println("Please enter a year:")
		fmt.Scan(&number)
		if number == 0 {
			fmt.Println("End of program")
			break;
		} else if leap(number) {
			fmt.Printf("%v is a leap year.\n", number)
		} else {
			fmt.Printf("%v is not a leap year.\n", number)
		}
	}
}

func leap(i int) bool {
	switch {
	case i%4 != 0:
		return false
	case i%100 != 0:
		return true
	case i%400 == 0:
		return true
	default:
		return false
	}
}
