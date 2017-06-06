package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

// WordCount is exported
func WordCount(s string) map[string]int {
	wcMap := map[string]int{}

	for _, word := range strings.Fields(s) {
		wcMap[word]++
	}
	//for _,word := range strings.Fields(s) {
	//	if _, ok := wcMap[word]; ok {
	//		wcMap[word]++
	//	} else {
	//		wcMap[word] = 1
	//	}
	//}
	return wcMap
}

func main() {
	wc.Test(WordCount)
}
