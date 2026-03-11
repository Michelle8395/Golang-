package main

import (
	"fmt"
)

func LongestWord(str string) string {
	longest := ""
	current := ""

	for i := 0; i < len(str); i++ {
		c := str[i]

		if c != ' ' {
			current += string(c) // add char to current word
		}

		if c == ' ' || i == len(str)-1 {
			// End of a word
			if len(current) > len(longest) {
				longest = current
			}
			current = "" // reset for next word
		}
	}

	return longest
}

func main() {
	s := "Go is an amazing programming language"
	fmt.Println("Longest word is:", LongestWord(s))
}