func numberOfMatches(n int) int {
    result := 0
    for n > 1 {
        if ((n&1) == 0) {
            result += n/2
            n = n/2
        } else {
            result += (n-1) / 2
            n = ((n-1) / 2) + 1
        }
    }
    
    return result
}