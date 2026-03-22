package piscine

func RepeatAlpha(s string) string {
	result := ""

	for i := 0; i < len(s); i++ {
		c := s[i]

		// Lowercase letters
		if c >= 'a' && c <= 'z' {
			count := int(c - 'a' + 1)
			for j := 0; j < count; j++ {
				result += string(c)
			}
		} else if c >= 'A' && c <= 'Z' { // Uppercase letters
			count := int(c - 'A' + 1)
			for j := 0; j < count; j++ {
				result += string(c)
			}
		} else {
			// Non-alphabetic characters remain unchanged
			result += string(c)
		}
	}

	return result
}