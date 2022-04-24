package main

import "fmt"

func main() {
    switch "Nasko" {
    case "Atanas":
        fmt.Println("Whatsup Atanas?")
    case "Vasilev":
        fmt.Println("Whatsup Vasilev?")
    case "Nakata":
        fmt.Println("Whatsup Nakata?")
    default:
        fmt.Println("Don't you have friends?")
    }
}
