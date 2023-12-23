func twoSum(nums []int, target int) []int {
	numIndex := make(map[int]int)
	for index, value := range nums {
		complement := target - value
		_, ok := numIndex[complement]
		if ok {
			return []int{numIndex[complement], index}
		}
		numIndex[value] = index
	}

	return nil
}