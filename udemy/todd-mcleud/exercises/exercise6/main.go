package main

import "fmt"

// Write a program that prints the numbers from 1 to 100.
// But for multiples of three print "Fizz" instead of the number
// and for the multiples of five print "Buzz".
//
// For numbers which are multiples of both three and five print "FizzBuzz".

func main() {
    for i := 1; i < 101; i++ {
        var fb string
        if i%3 == 0 {
            fb += "Fizz"
        }
        if i%5 == 0 {
            fb += "Buzz"
        }
        switch {
        case len(fb) > 0:
            fmt.Println(fb)
        default:
            fmt.Println(i)
        }
    }
}
