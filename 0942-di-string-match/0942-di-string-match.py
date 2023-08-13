class Solution:
    def diStringMatch(self, s: str) -> List[int]:
        result = []
        smallest = 0
        largest = len(s)
        for char in s:
            if char == "I":
                result.append(smallest)
                smallest += 1
            else:
                result.append(largest)
                largest -= 1
        result.append(largest) if s[-1] == "D" else result.append(smallest)
        return result