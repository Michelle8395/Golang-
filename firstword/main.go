package main 

func FirstWord(s string) string {

	word := ""
	i := 0

	// skip leading spaces
	for i < len(s) && s[i] == ' ' {
		i++
	}

	// collect first word
	for i < len(s) && s[i] != ' ' {
		word += string(s[i])
		i++
	}

	return word + "\n"
}