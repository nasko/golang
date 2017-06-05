package main

import "fmt"

/*
    Write a function which takes an integer. The function will have two returns.
    The first return should be the argument divided by 2.
    The second return should be a bool that let’s us know whether or not the argument was even.
    For example:
    half(1) returns (0, false)
    half(2) returns (1, true)
*/

func halfb(i int) (float64, bool) {
    return float64(i)/2, i%2 == 0
}

func main() {

    fmt.Println(halfb(1))
    fmt.Println(halfb(2))
    fmt.Println(halfb(5))
}
