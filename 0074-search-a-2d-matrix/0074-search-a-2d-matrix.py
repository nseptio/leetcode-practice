class Solution:
    def searchMatrix(self, matrix: List[List[int]], target: int) -> bool:
        if not matrix:
            return False

        rows = len(matrix)
        cols = len(matrix[0])

        left = 0
        right = rows * cols - 1

        while left <= right:
            mid_point = left + (right - left) // 2
            mid_point_row = mid_point // cols
            mid_point_col = mid_point % cols
            mid_point_element = matrix[mid_point_row][mid_point_col]

            if target == mid_point_element:
                return True
            elif target < mid_point_element:
                right = mid_point - 1
            else:
                left = mid_point + 1

        return False