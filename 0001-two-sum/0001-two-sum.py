class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        num_index = {}
        for i in range(0, len(nums)):
            complement = target - nums[i]
            if complement in num_index:
                return [num_index[complement], i]
            num_index[nums[i]] = i
        
        return None
                