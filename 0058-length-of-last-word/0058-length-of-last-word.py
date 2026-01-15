class Solution:
    def lengthOfLastWord(self, s: str) -> int:
        stripped = s.strip()
        word_list = stripped.split()
        return len( word_list[-1])
        