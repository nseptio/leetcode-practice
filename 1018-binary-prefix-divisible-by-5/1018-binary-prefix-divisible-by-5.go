func prefixesDivBy5(nums []int) []bool {
    result := make([]bool, len(nums))
    current := 0

    for i, num := range nums {
        current = (current*2 + num) % 5
        result[i] = (current == 0)
    }

    return result
}