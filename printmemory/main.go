package piscine

func PrintMemory(arr [10]byte) string {
    hex := "0123456789abcdef"
    result := ""

    // Hexadecimal part
    for i := 0; i < 10; i++ {
        b := arr[i]
        result += string(hex[b/16]) + string(hex[b%16])

        if i == 3 || i == 7 || i == 9 {
            result += "\n"
        } else {
            result += " "
        }
    }

    // ASCII part
    for i := 0; i < 10; i++ {
        if arr[i] >= 32 && arr[i] <= 126 {
            result += string(arr[i])
        } else {
            result += "."
        }
    }

    result += "\n"
    return result
}