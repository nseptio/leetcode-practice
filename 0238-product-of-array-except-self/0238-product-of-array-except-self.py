class Solution:
    def productExceptSelf(self, nums: List[int]) -> List[int]:
        result = []
        result.append(1)
        
        for i in range(len(nums) - 1):
            prefix = result[i] * nums[i]
            result.append(prefix)

        suffix = 1
        for j in range(len(nums) - 1, -1, -1):
            result[j] = suffix * result[j]
            suffix *= nums[j]
        
        return result