func restoreString(s string, indices []int) string {
    chars := make([]rune, len(s))
    for i, ch := range s {
        chars[indices[i]] = ch
    }
    
    return string(chars)
}