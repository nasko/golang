package main

import "fmt"

/**
    Write a function with one variadic parameter that finds the greatest number in a list of numbers
 */
func main() {
    series := []int{5,6,67,32,89}
    fmt.Println(max(series...))

    fmt.Println(max(5,3,2,7,4))
}

func max(numbers ...int) int {
    var max int

    for _,val := range numbers {
        switch {
        case val > max:
            max = val
        }
    }
    return max
}
