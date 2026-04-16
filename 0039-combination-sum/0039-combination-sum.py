class Solution:
    def combinationSum(self, nums: List[int], target: int) -> List[List[int]]:
        res = []

        def dfs(i, total, curr_nums):
            if total == target:
                res.append(list(curr_nums))
                return

            if total > target or i >= len(nums):
                return

            curr_nums.append(nums[i])
            dfs(i, total + curr_nums[-1], curr_nums)

            curr_nums.pop()
            dfs(i + 1, total, curr_nums)

        dfs(0, 0, [])
        return list(res)