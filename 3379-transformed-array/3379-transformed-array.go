func constructTransformedArray(nums []int) []int {
    result := make([]int, len(nums))

    for i := 0; i < len(nums); i++ {
        index := nums[i]
        if index > 0 {
            index += i
            for index > len(nums) - 1 {
                index -= len(nums)
            }
        } else if index < 0 {
            index += i
            for index < 0 {
                index += len(nums)
            }
        } else if index == 0 {
            result[i] = nums[i]
            continue
        }

        result[i] = nums[index]
    }

    return result
}
