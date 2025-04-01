func checkOnesSegment(s string) bool {
    total := 0

    isOneSeq := false
    for i := 0; i < len(s); i++ {
        if isOneSeq {
            if string(s[i]) == "0" {
                isOneSeq = false
            }
        } else {
            if string(s[i]) == "1" {
                total += 1
                isOneSeq = true
            }
        }
    }

    return total == 1
}