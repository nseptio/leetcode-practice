import (
    "strings"
    "strconv"
)

func areNumbersAscending(s string) bool {
    stringSlice := strings.Split(s, " ")
    lastNumber := -1
    currentNumber := -1
    for i := 0; i < len(stringSlice); i++ {
        v, err := strconv.Atoi(stringSlice[i])
        if err != nil {
            continue
        }

        if lastNumber == -1 {
            lastNumber = v
            continue
        }

        currentNumber = v
        
        if currentNumber <= lastNumber {
            return false
        }
        
        lastNumber = currentNumber
    }

    return true
}