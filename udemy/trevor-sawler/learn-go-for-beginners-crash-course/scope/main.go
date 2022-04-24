package main

import (
	"myapp/packageone"
)

var myVar = "myVar"

func main() {
	// variables
	// declare a package lever variable for the main
	// package named myVar

	// declare a block variable for the main func called bockVar
	var blockVar = "blockVar"

	// declare a package level var in the packageone package named PackageVar

	// create an exported function in packageone called PrintMe

	// in the main function print out the values of myVar, blockVar, and PackageVar on one line
	// using the PrintMe func in packageone
	packageone.PrintMe(myVar, blockVar)

}
