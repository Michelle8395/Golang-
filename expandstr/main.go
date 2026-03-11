package main

import "os"
import "fmt"

func main(){
	if len(os.Args) != 2{

		return
	}

	s := os.Args[1]
	printedWord := false
	result := ""
	i := 0

	//loop thru the str

	for i < len(s){

		//skip trailing whitespaces

		for i < len(s) && s[i] == ' '{
			i ++
		}

		if i > len(s){
			break
		}

		if printedWord {
			result += "   "
		}

		for i < len(s) && s[i] != ' '{
			result += string(s[i])

			i++
		}
		printedWord = true


	}

	fmt.Println(result)


}