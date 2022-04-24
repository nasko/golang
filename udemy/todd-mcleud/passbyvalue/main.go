package main

import "fmt"

func changeme(z *int) {
    fmt.Println(z)
    fmt.Println(*z)
    *z = 24
    fmt.Println(z)
    fmt.Println(*z)
}

func main() {
    age := 19
    fmt.Println(&age)
    changeme(&age)
    fmt.Println(&age)
    fmt.Println(age)
}
