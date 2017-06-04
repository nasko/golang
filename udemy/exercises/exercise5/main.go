package main

import "fmt"

// Print all of the even numbers between 0 and 100

func main() {

    for i := 0; i < 101; i++ {
        if 0 == i & 1 {
            fmt.Println(i)
        }
    }
}
