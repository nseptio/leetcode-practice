class Solution:
    def longestCommonPrefix(self, strs: List[str]) -> str:
        sorted_list = sorted(strs)
        result = ""
        first_str = sorted_list[0]
        last_str = sorted_list[-1]
        for i in range(0, len(first_str)):
            if first_str[i] != last_str[i]:
                break
            result += first_str[i]
        return result