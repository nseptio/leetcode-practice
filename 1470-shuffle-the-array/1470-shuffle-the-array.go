func shuffle(nums []int, n int) []int {
    result := make([]int, 2*n)
    idx := 0
    for i := 0; i < n; i++ {
        result[idx] = nums[i]
        result[idx+1] = nums[n+i]
        idx += 2
    }
    return result
}