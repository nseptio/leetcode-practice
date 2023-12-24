func twoSum(nums []int, target int) []int {
	numIndex := make(map[int]int)
	for index, value := range nums {
		complement := target - value
		mapValue, ok := numIndex[complement]
		if ok {
			return []int{mapValue, index}
		}
		numIndex[value] = index
	}

	return nil
}