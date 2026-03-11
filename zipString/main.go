package main

import "fmt"

func ZipString(s string) string {

	result := ""
	count := 1
	currentChar := s[0]

	for i := 1; i < len(s); i++ {
		if s[i] == currentChar {
			count++
		} else {
			result += string(count+'0') + string(currentChar)

			count = 1
			currentChar = s[i]
		}

	}

	result += string(count+'0') + string(currentChar)
	return result

}

func main() {
	fmt.Println(ZipString("YouuungFellllas"))
	fmt.Println(ZipString("Thee quuick browwn fox juumps over the laaazy dog"))
	fmt.Println(ZipString("Helloo Therre!"))
}
