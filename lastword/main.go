package main 

func LastWord(s string) string {

	word := ""
	i := len(s) - 1

	// skip trailing spaces
	for i >= 0 && s[i] == ' ' {
		i--
	}

	// collect last word
	for i >= 0 && s[i] != ' ' {
		word = string(s[i]) + word
		i--
	}

	return word + "\n"
}