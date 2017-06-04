package main

import "fmt"

func main() {
    for i:=5000; i<5100; i++ {
        fmt.Printf("%v\t%v\t%v\t%v\n", i, string(i), []byte(string(i)),byte(i))
    }
}
