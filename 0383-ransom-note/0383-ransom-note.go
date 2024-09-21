func canConstruct(ransomNote string, magazine string) bool {
    ransomMap := make(map[rune]int)
    magazineMap := make(map[rune]int)
    for _, ch := range ransomNote {
        ransomMap[ch]++
    }
    
    for _, ch := range magazine {
        magazineMap[ch]++
    }
    
    for k, v := range ransomMap {
        if v > magazineMap[k] {
            return false
        }
    }
    
    return true
}