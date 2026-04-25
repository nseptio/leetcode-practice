class Solution:
    def minCostClimbingStairs(self, cost: List[int]) -> int:
        one, two = cost[-1], cost[-2]
        for i in range(len(cost) - 3, -1, -1):
            temp = min(cost[i] + one, cost[i] + two)
            one = two
            two = temp
        
        return min(one, two)
            