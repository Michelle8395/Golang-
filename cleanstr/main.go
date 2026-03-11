package main 
import "os"
import "fmt"

func main(){
    if len(os.Args) != 2{
        return
    }

    s := os.Args[1]
    i := 0
    wordPrinted := false
    result := ""

    for i < len(s) {

    // skip spaces
    for i < len(s) && s[i] == ' ' {
        i++
    }

    if i >= len(s) {
        break
    }

    // add space before next word
    if wordPrinted {
        result += " "
    }

    // read word
    for i < len(s) && s[i] != ' ' {
        result += string(s[i])
        i++
    }

    wordPrinted = true
}

fmt.Println(result)

}