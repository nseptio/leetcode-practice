class Solution:
    def maxLengthBetweenEqualCharacters(self, s: str) -> int:
        indexes = {}
        for i in range(0, len(s)):
            if s[i] not in indexes:
                indexes[s[i]] = [i]

            idx = s.find(s[i], i + 1)
            if idx != -1:
                indexes[s[i]].append(idx)

        max_len = -1
        for key, value in indexes.items():
            if len(value) >= 2:
                diff = value[-1] - value[0] - 1
                max_len = diff if diff > max_len else max_len

        return max_len
            
            