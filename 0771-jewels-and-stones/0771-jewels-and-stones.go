import "slices"

func numJewelsInStones(jewels string, stones string) int {
    jewelsRune := []rune(jewels)

    result := 0
    for _, stone := range stones {
        if slices.Contains(jewelsRune, stone) {
            result += 1
        }
    }
    
    return result
}