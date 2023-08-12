class Solution:
    def minimumSum(self, num: int) -> int:
        num_str = list(str(num))
        num_str.sort()
        return int(num_str[0]+num_str[3]) + int(num_str[1] + num_str[2])