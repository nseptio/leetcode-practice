class Solution:
    def strStr(self, haystack: str, needle: str) -> int:
        if needle not in haystack:
            return -1
        
        for i in range(len(haystack)):
            if i + len(needle) > len(haystack):
                return -1
            elif needle == haystack[i:i+len(needle)]:
                return i
        
        return -1