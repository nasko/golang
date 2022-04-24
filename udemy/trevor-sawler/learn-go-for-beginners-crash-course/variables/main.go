package main

import "fmt"

func main() {
	makeArrays()
	makeSlices()
	makeMaps()
}

func makeArrays() {
	var animalsArray [3]string
	animalsArray[0] = "cat"
	animalsArray[1] = "dog"
	animalsArray[2] = "fish"

	peopleArray := [3]string{
		"John",
		"Jane",
		"Paul",
	}
	fmt.Println("peopleArray:", peopleArray)

	// If you omit the size of the array, an array just big enough to hold the initialization is created. Therefore, if you write −
	var balanceArray = []float32{1000.0, 2.0, 3.4, 7.0, 50.0}
	fmt.Println("balanceArray:", balanceArray)

	// Iterate over arrays

	// for .. range assigning both index and value
	for index, value := range peopleArray {
		fmt.Println(index, value)
	}
	// for .. range assigning just the index
	for index := range peopleArray {
		fmt.Println(index, peopleArray[index])
	}
	// for .. range assigning just the value
	for _, value := range peopleArray { // escaping index by _ keyword
		fmt.Println(value)
	}

	// for .. len(array)
	for i := 0; i < len(animalsArray); i++ {
		fmt.Printf("%d th element of animalsArray is %s\n", i, animalsArray[i])
	}
}

func makeSlices() {
	// To define a slice, you can declare it as an array without specifying its size
	var animalsSlice []string
	animalsSlice = append(animalsSlice, "dog")
	animalsSlice = append(animalsSlice, "fish")
	animalsSlice = append(animalsSlice, "cat")
	animalsSlice = append(animalsSlice, "horse")
	fmt.Println("animalsSlice:", animalsSlice)

	// Alternatively, you can use make function to create a slice.
	numbersSlice := make([]int, 0)
	numbersSlice = append(numbersSlice, 1)
	numbersSlice = append(numbersSlice, 2)
	fmt.Println("numbersSlice:", numbersSlice)

	// Declare and initialize a slice
	daysSlice := []string{
		"Monday",
		"Tuesday",
		"Wednesday",
	}
	fmt.Println("daysSlice:", daysSlice)

	// Create a slice by re-slicing an existing slice
	numbers := []int{1, 2, 3, 4, 5}

	//Both start and end
	num1 := numbers[2:4]
	fmt.Println("Both start and end")
	fmt.Printf("num1=%v\n", num1)
	fmt.Printf("length=%d\n", len(num1))
	fmt.Printf("capacity=%d\n", cap(num1))

	//Only start
	num2 := numbers[2:]
	fmt.Println("\nOnly start")
	fmt.Printf("num1=%v\n", num2)
	fmt.Printf("length=%d\n", len(num2))
	fmt.Printf("capacity=%d\n", cap(num2))

	//Only end
	num3 := numbers[:3]
	fmt.Println("\nOnly end")
	fmt.Printf("num1=%v\n", num3)
	fmt.Printf("length=%d\n", len(num3))
	fmt.Printf("capacity=%d\n", cap(num3))

	//None
	num4 := numbers[:]
	fmt.Println("\nNone")
	fmt.Printf("num1=%v\n", num4)
	fmt.Printf("length=%d\n", len(num4))
	fmt.Printf("capacity=%d\n", cap(num4))

	// Iterate over slices

	// Assigning both the index and the value
	for i, x := range animalsSlice {
		fmt.Printf("animalsSlice[%d]: %s\n", i, x)
	}
	// Assigning just the index
	for i := range animalsSlice {
		fmt.Printf("animalsSlice[%d]: %s\n", i, animalsSlice[i])
	}
	// Assigning just the value
	for _, x := range animalsSlice {
		fmt.Println(x)
	}

}

func makeMaps() {

	// Declare using an empty map literal
	statesMap := map[string]string{}
	statesMap["NY"] = "New York"
	statesMap["CA"] = "California"
	fmt.Println("statesMap:", statesMap)

	// Declare and initialize, using the map literal with some key-values initilized
	employeeSalary := map[string]int{
		"John": 1000,
		"Sam":  2000,
	}
	employeeSalary["Tom"] = 2000
	fmt.Println("employeeSalary:", employeeSalary)

	// Declaring using var and then adding key-value pairs won't work:

	// var monthlyIncome map[string]int
	// panic: assignment to entry in nil map
	// monthlyIncome["January"] = 2000

	// Create a map using the builtin function make()
	// It returns an initialized map
	// Hence key-value pairs can be added to it
	intMap := make(map[string]int)

	intMap["one"] = 1
	intMap["two"] = 2
	intMap["three"] = 3
	intMap["four"] = 4
	intMap["five"] = 5

	// Iterate over a map
	for key, value := range statesMap {
		fmt.Println(key, value)
	}
	// Or assigning just the key
	for key := range intMap {
		fmt.Println(key, intMap[key])
	}
	// Or assigning just the value
	for _, value := range statesMap {
		fmt.Println(value)
	}

	// Check to see if an element with a certain key exists in the map
	value, ok := employeeSalary["Nobody"]
	if ok {
		fmt.Println("Yes the element exists in the map:", value)
	} else {
		// returns the corresponding zero-value for the relevant type - 0 in this case
		fmt.Println("Non-existing value:", value) // 0
	}
}
