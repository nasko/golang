package main

import "fmt"

/**
Write a function, foo, which can be called in all of these ways:

func main() {
	foo(1, 2)
	foo(1, 2, 3)
	aSlice := []int{1, 2, 3, 4}
	foo(aSlice...)
	foo()
}
*/

func main() {
    foo(1, 2)
    foo(1, 2, 3)
    aSlice := []int{1, 2, 3, 4}
    foo(aSlice...)
    foo()
}

func foo(numbers ...int) {
    var max int

    for _,val := range numbers {
        switch {
        case val > max:
            max = val
        }
    }
    fmt.Println(max)
}
