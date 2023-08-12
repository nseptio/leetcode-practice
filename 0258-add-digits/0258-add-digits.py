from functools import reduce

class Solution:
    def addDigits(self, num: int) -> int:
        digits = [int(digit) for digit in str(num)]
        sum_result = reduce(lambda x, y: x + y, digits)
        while (sum_result // 10 > 0):
            digits = [int(digit) for digit in str(sum_result)]
            sum_result = reduce(lambda x, y: x + y, digits)
        return sum_result
