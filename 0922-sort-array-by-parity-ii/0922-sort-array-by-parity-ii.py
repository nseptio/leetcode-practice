class Solution:
    def sortArrayByParityII(self, nums: List[int]) -> List[int]:
        odd_list = sorted([i for i in nums if i % 2 == 1])
        even_list = sorted([i for i in nums if i % 2 == 0])
        result = []
        i, j = 0, 0
        for k in range(0, len(odd_list) + len(even_list)):
            if k % 2 == 0:
                result.append(even_list[i])
                i += 1
            else:
                result.append(odd_list[j])
                j += 1
                
        return result