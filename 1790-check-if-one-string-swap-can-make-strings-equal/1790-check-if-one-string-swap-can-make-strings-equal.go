func areAlmostEqual(s1 string, s2 string) bool {
    diff := 0
    for i := 0; i < len(s1); i++ {
        if s1[i] != s2[i] {
            diff++
        }
    }

    if diff > 2 {
        return false
    }

    return true
}