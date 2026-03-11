package piscine

func WeAreUnique(str1, str2 string) int {

    if str1 == "" && str2 == "" {
        return -1
    }

    count := 0
    used := "" // store counted characters

    // helper function to check if char is in string
    isIn := func(s string, c byte) bool {
        for i := 0; i < len(s); i++ {
            if s[i] == c {
                return true
            }
        }
        return false
    }

    // check str1
    for i := 0; i < len(str1); i++ {
        c := str1[i]
        if !isIn(str2, c) && !isIn(used, c) {
            count++
            used += string(c)
        }
    }

    // check str2
    for i := 0; i < len(str2); i++ {
        c := str2[i]
        if !isIn(str1, c) && !isIn(used, c) {
            count++
            used += string(c)
        }
    }

    return count
}