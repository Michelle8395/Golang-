package main

import "os"
import "fmt"

func main(){
	if len(os.Args) != 2{
		return
	}

	s := os.Args[1]
	printedWord := false
	i := 0
	word := ""
	result := ""

	for i < len(s){

		 //skip trailing spaces

		 for i < len(s) && s[i] == ' '{
			i++
		 }

		 //read word

		 for i < len(s) && s[i] != ' '{
			word += string(s[i])
			i++
		 }

		 if word != ""{
			if printedWord{
				result += " "
			}
			result += word
			printedWord = true
			word = ""

		
		 }
}
	fmt.Println(result)

}