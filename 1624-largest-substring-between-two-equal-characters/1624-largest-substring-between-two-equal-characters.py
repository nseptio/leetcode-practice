class Solution:
    def maxLengthBetweenEqualCharacters(self, s: str) -> int:
        indexes = {}
        for i in range(0, len(s)):
            if s[i] not in indexes:
                indexes[s[i]] = [i]

            idx = s.find(s[i], i + 1)
            if idx != -1:
                indexes[s[i]].append(idx)

        max_len = 0
        is_available = False
        for key, value in indexes.items():
            if len(value) >= 2:
                is_available = True
                diff = value[-1] - value[0] - 1
                max_len = diff if diff > max_len else max_len

        return max_len if is_available else -1
            
            