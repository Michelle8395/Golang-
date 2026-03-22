package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 4 {
		return // do nothing if argument count is not 3
	}

	input := os.Args[1]
	oldChar := os.Args[2]
	newChar := os.Args[3]

	// Make sure oldChar and newChar are single characters
	if len(oldChar) != 1 || len(newChar) != 1 {
		return
	}

	result := ""

	for i := 0; i < len(input); i++ {
		if string(input[i]) == oldChar {
			result += newChar
		} else {
			result += string(input[i])
		}
	}

	fmt.Println(result)
}