class Solution:
    def maxArea(self, height: List[int]) -> int:
        left = 0
        right = len(height) - 1
        max_area = -1
        for i in range(left, len(height)):
            h = min(height[left], height[right])
            area = h * (right - left)
            max_area = area if area > max_area else max_area

            if abs(left - right) != 1:
                if height[left] < height[right]:
                    left += 1
                else:
                    right -= 1

        return max_area