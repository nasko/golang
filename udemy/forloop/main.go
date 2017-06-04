package main

import "fmt"

func main() {
    for i := 0; i < 10; i++ {
        fmt.Println(i)
    }

    fmt.Println("Second loop:")
    var i int
    for i < 20 {
        fmt.Println(i)
        i++
    }

    fmt.Println("Third loop:")
    var x int
    for {
        x++

        // skip  odd numbers
        if bool(x & 1) {
            continue
        }

        fmt.Println(x)

        if x == 20 {
            break;
        }
    }



}
