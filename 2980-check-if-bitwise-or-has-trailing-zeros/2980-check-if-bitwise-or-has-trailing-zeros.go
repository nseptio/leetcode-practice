func hasTrailingZeros(nums []int) bool {
    
    check := make([]int, 0, 2)
    for i := 0; i < len(nums); i++ {
        if nums[i] & 1 == 0 {
            check = append(check, i)
        }
        if len(check) == 2 {
            return true
        }
    }
    
    return len(check) >= 2
}