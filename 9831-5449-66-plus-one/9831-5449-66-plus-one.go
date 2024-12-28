func plusOne(digits []int) []int {
    lastIndex := len(digits) - 1
    if digits[lastIndex] != 9 {
        digits[lastIndex] += 1
        return digits
    }

    for i := lastIndex; i >= 0; i-- {
        if i == 0 {
            if digits[i] != 9 {
                digits[i] += 1
            } else {
                digits[i] = 1
                digits = append(digits, 0)
            }
            break
        }

        if digits[i] != 9 {
            digits[i] += 1
            break
        }
        digits[i] = 0
    }

    return digits
}