import "strconv"

func getLucky(s string, k int) int {
    str := ""
    for _, ch := range s {
        str += strconv.Itoa(int(ch) - 96)
    }
    
    result := 0
    for i := 0; i < k; i++ {
        for _, v := range str {
            result += int(v) - 48
        }
        str = strconv.Itoa(result)
        if (i != k - 1) {
            result = 0
        }
    }
    
    return result
}