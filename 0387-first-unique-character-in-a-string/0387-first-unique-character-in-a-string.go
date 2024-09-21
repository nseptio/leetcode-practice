func firstUniqChar(s string) int {
    charMap := make(map[rune]int)
    for _, ch := range s {
        charMap[ch]++
    }
    
    for i, ch := range s {
        if charMap[ch] == 1 {
            return i
        }
    }
    
    return -1
}