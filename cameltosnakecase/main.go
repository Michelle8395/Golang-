package main

import "fmt"

func CamelToSnakeCase(s string) string{

	result := ""

	if s == ""{
		return ""
	}

	for i := 0; i < len(s); i++{

		//check for non alphabets

		if !(s[i] >= 'A' && s[i] <= 'Z') && !(s[i] >= 'a' && s[i] <= 'z'){
			return s
		}

		//check for concurrent recurrence

		if i > 0 && s[i] >= 'A' && s[i] <= 'Z' && s[i-1] >= 'A' && s[i-1] <= 'Z'{
			return s
		}

		//check last uppercase

		if i == len(s)-1 && s[i] >= 'A' && s[i] <= 'Z'{
			return s
		}

		if s[i] >= 'A' && s[i] <= 'Z'{
			if i != 0{
				result += "_"
			}
			result += string(s[i])
		}else{
			result += string(s[i])
				
			}
	
	}
	return result
}

func main() {
	fmt.Println(CamelToSnakeCase("HelloWorld"))
	fmt.Println(CamelToSnakeCase("helloWorld"))
	fmt.Println(CamelToSnakeCase("camelCase"))
	fmt.Println(CamelToSnakeCase("CAMELtoSnackCASE"))
	fmt.Println(CamelToSnakeCase("camelToSnakeCase"))
	fmt.Println(CamelToSnakeCase("hey2"))
}
