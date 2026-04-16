class Solution:
    def combinationSum2(self, candidates: List[int], target: int) -> List[List[int]]:
        res = []
        candidates.sort()

        def dfs(i: int, total: int, curr: List[int]):
            if total == target:
                res.append(curr.copy())
                return
            if i >= len(candidates) or total > target:
                return
            
            curr.append(candidates[i])
            dfs(i + 1, total + curr[-1], curr)

            curr.pop()
            
            while i+1 < len(candidates) and candidates[i] == candidates[i+1]:
                i +=1
            dfs(i + 1, total , curr)
        
        dfs(0, 0, [])
        return res