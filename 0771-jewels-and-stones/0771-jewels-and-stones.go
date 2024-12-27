func numJewelsInStones(jewels string, stones string) int {
    jewelSet := make(map[rune]bool)
    for _, ch := range jewels {
        _, ok := jewelSet[ch]
        if !ok {
            jewelSet[ch] = true
        }
    }

    result := 0
    for _, stone := range stones {
        _, ok := jewelSet[stone]
        if ok {
            result += 1
        }
    }
    
    return result
}