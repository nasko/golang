package main

import (
	"strings"
	"fmt"
)

/*
Bob

Bob is a lackadaisical teenager. In conversation, his responses are very limited.
Bob answers 'Sure.' if you ask him a question.
He answers 'Whoa, chill out!' if you yell at him.
He says 'Fine. Be that way!' if you address him without actually saying anything.
He answers 'Whatever.' to anything else.

To test, run go test in the exercise folder
    - the Hey() method will be tested against all the messages in cases_test.go
*/

// Hey is exported
func Hey(s string) string {
	var hasLowerCase, hasUpperCase bool

	isCaps := func(rSlice []rune) bool {
		for _, r := range rSlice {
			dec := int(r)
			if dec >= 65 && dec <= 90 {
				hasUpperCase = true
			}
			if dec >= 97 && dec <= 122 {
				hasLowerCase = true
			}
		}
		if !hasLowerCase && hasUpperCase {
			return true
		}
		return false
	}

	s = strings.TrimSpace(s)

	if len(s) == 0 {
		return "Fine. Be that way!"
	}

	rSlice := []rune(s)

	if isCaps(rSlice) {
		return "Whoa, chill out!"
	}

	lastChar := string(rSlice[len(rSlice)-1])
	if "?" == lastChar {
		return "Sure."
	}

	return "Whatever."
}

func main() {
	fmt.Println(Hey("Tom-ay-to, tom-aaaah-to."))
}
