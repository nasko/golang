package main

import (
	"fmt"

	"github.com/eiannone/keyboard"
)

// Create a channel that only accepts the rune type
var keyPressChan chan rune

func main() {
	keyPressChan = make(chan rune)

	// Starts the listenForKeyPress() function in its own async go-routine
	go listenForKeyPress()

	fmt.Println("Press any key, or q to quit")
	_ = keyboard.Open()
	defer func() {
		keyboard.Close()
	}()

	for {
		char, _, _ := keyboard.GetSingleKey()
		if char == 'q' || char == 'Q' {
			break
		}

		// Passes the rune to the channel
		keyPressChan <- char
	}
}

func listenForKeyPress() {
	for {
		// Receives the rune from the channel
		key := <-keyPressChan
		fmt.Println("You pressed", string(key))
	}
}
