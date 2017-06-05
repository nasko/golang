package main

import (
    "fmt"
    "strconv"
)

// Here's how to get the digits of an integer

func main() {
    // Generally range iterates over the bytes of a string, but if you want to get a slice
    // consisting of those bytes:
    //chars := []byte(strconv.Itoa(8362374835))

    for _,c := range strconv.Itoa(836237445456353453) {
        fmt.Println(string(c))
    }
}
