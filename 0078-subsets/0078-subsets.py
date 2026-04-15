class Solution:
    def subsets(self, nums: List[int]) -> List[List[int]]:
        index = 0
        res = []

        def search_subsets(i, curr):
            if i == len(nums):
                res.append(curr)
                return
            
            curr_added = curr + [nums[i]]
            i += 1
            search_subsets(i, curr_added)
            search_subsets(i, curr)
        
        search_subsets(index, [])
        return res