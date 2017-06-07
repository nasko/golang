package main

import (
	"sort"
	"fmt"
)

type people []string

func main() {
	studyGroup := people{
		"Zeno",
		"John",
		"Al",
		"Jenny",
	}

	// original
	fmt.Println(studyGroup)

	// ascending
	sort.Sort(sort.StringSlice(studyGroup))
	fmt.Println(studyGroup)

	// descending
	sort.Sort(sort.Reverse(sort.StringSlice(studyGroup)))
	fmt.Println(studyGroup)
}
