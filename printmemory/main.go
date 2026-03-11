package main

import "fmt"

func PrintMemory(arr [10]byte) {
	hexChars := "0123456789abcdef"

	// Print hex, 4 bytes per line
	for i := 0; i < len(arr); i++ {
		b := arr[i]
		fmt.Printf("%c%c", hexChars[b/16], hexChars[b%16])
		if (i+1)%4 == 0 || i == len(arr)-1 {
			fmt.Println()
		} else {
			fmt.Print(" ")
		}
	}

	// Print ASCII characters
	for i := 0; i < len(arr); i++ {
		b := arr[i]
		if b >= 32 && b <= 126 {
			fmt.Printf("%c", b)
		} else {
			fmt.Print(".")
		}
	}
	fmt.Println()
}