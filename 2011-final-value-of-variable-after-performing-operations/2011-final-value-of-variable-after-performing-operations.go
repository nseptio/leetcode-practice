import "strings"

func finalValueAfterOperations(operations []string) int {
    x := 0
    for _, op := range operations {
        if strings.ContainsRune(op, '+' ) {
            x += 1
        } else {
            x -= 1
        }
    }
    
    return x
}