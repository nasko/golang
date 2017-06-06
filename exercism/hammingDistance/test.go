package main

import "fmt"

func main() {
	name := "атанас"

	rName := []rune(name)
	fmt.Println([]rune(name))

	for _,i := range rName {
		fmt.Println(string(i))
	}

	//for _,char := range name {
	//	fmt.Println(string(char))
	//}
	//for i,_ := range name {
	//	fmt.Println([]rune(name)[i])
	//}
	//chars := []rune(name)

	//for i,_ := range name {
	//	fmt.Printf("%v\n", string([]rune(name)[i]))
	//}
}
