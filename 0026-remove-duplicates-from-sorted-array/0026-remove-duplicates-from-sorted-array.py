class Solution:
    def removeDuplicates(self, nums: List[int]) -> int:
        num_set = set()
        unique_len = set(nums)
        i = 0
        while i < len(unique_len):
            if nums[i] in num_set:
                popped = nums.pop(i)
                nums.append(popped)
            else:
                num_set.add(nums[i])
                i += 1
        
        return len(num_set)