import "math"

func titleToNumber(columnTitle string) int {
    result := 0
    for i:= 0; i < len(columnTitle); i++ {
        value := int(columnTitle[i] - 'A') + 1
        result += value * int(math.Pow(26, float64(len(columnTitle) - i - 1)))
    }
    return result
}