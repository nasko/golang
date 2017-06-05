package main

import "fmt"

func main() {
    data := []float64{2, 4, 6}

    a := average(data...)
    b := average(2, 4)
    fmt.Println(a)
    fmt.Println(b)
}
func average(sf ...float64) float64 {
    var total float64

    for _, v := range sf {
        total += v
    }
    return total / float64(len(sf))
}
