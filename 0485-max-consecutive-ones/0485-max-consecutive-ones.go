func findMaxConsecutiveOnes(nums []int) int {
    size := 0
    result := 0
    for i := 0; i < len(nums); i++ {
        if nums[i] == 1 {
            size++
        }
        
        if size > result {
            result = size
        }
        
        if nums[i] == 0 {
            size = 0
        }
    }
    return result
}