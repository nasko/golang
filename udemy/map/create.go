package main

import "fmt"

func main() {
	// variant 1
	var phoneBook = map[string]string{}
	phoneBook["Atanas Vasilev"] = "+359888263706"
	phoneBook["Атанас служ."] = "+359894344450"
	fmt.Println(phoneBook)

	// variant 2
	phones := map[string]string{
		"nasko": "088263706",
		"Атанас": "0894344450",
	}
	fmt.Println(phones)
	fmt.Println(phones["Атанас"])

	// variant 3
	var recipes = make(map[string]string)
	recipes["pancakes"] = "Lorem ipsum"
	recipes["баница"] = "Lorem ipsum"
	fmt.Println(recipes)
	fmt.Println(recipes["баница"])

	_,existsBanitza := recipes["баница"]

	if existsBanitza {
		delete(recipes, "баница")
	}
	fmt.Println(recipes)

	if _,exists := recipes["pancakes"]; exists {
		delete(recipes, "pancakes")
	}
	fmt.Println("Recipes: ", recipes)
}
