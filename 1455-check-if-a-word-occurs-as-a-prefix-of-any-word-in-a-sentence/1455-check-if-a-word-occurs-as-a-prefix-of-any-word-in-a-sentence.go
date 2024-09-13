import (
    "strings"
)

func isPrefixOfWord(sentence string, searchWord string) int {
    fields := strings.Fields(sentence)
    for i, word := range fields {
        if (strings.HasPrefix(word, searchWord)) {
            return i+1
        }
    }
    
    return -1
}