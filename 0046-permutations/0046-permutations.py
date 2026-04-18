class Solution:
    def permute(self, nums: List[int]) -> List[List[int]]:
        res, curr = [], []

        def dfs():
            if len(curr) == len(nums):
                res.append(curr.copy())
                return
            
            for i in nums:
                if i not in curr:
                    curr.append(i)
                    dfs()
                    curr.pop()
        
        dfs()
        return res
        