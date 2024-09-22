func hasTrailingZeros(nums []int) bool {
    
    check := make([]int, 0, 2)
    for i := 0; i < len(nums); i++ {
        if nums[i] & 1 == 0 {
            check = append(check, i)
        }
    }
    
    return len(check) >= 2
}