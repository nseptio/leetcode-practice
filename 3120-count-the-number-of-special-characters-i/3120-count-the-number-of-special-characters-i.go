import "unicode"

func numberOfSpecialChars(word string) int {

    upperChar := make(map[rune]bool, 0)
    lowerChar := make(map[rune]bool, 0)
    result := 0
    for _, r := range word {
        if unicode.IsLower(r) {
            lowerChar[r] = true
        } else {
            upperChar[r] = true
        }
    }
    
    for k, _ := range lowerChar {
        _, ok := upperChar[unicode.ToUpper(k)]
        
        if ok {
            result+=1
        }
    }
    
    return result
}