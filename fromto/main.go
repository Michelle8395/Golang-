package main

import "fmt"

func FromTo(from int, to int) string {
	if from < 0 || from > 99 || to < 0 || to > 99 {
		return "Invalid\n"
	}

	result := ""
	step := 1
	if from > to {
		step = -1
	}

	for n := from; n != to; n += step {
		tens := n / 10
		ones := n % 10
		result += string('0'+tens) + string('0'+ones) + ", "
	}

	tens := to / 10
	ones := to % 10
	result += string('0'+tens) + string('0'+ones)

	return result + "\n"
}

func main() {
	fmt.Print(FromTo(1, 10))
	fmt.Print(FromTo(10, 1))
	fmt.Print(FromTo(10, 10))
	fmt.Print(FromTo(100, 10))
}