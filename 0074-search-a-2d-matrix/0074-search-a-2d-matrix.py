class Solution:
    def searchMatrix(self, matrix: List[List[int]], target: int) -> bool:
        m = len(matrix)
        n = len(matrix[0])

        s = [0, 0]
        e = [m - 1, n - 1]
        while s[0] <= e[0]:
            if s[0] == e[0] and s[1] > e[1]:
                break
            mid_index = [e[0] + s[0], e[1] + s[1]]
            mid_index_len = (n*mid_index[0] + mid_index[1]) // 2
            mid_index = [mid_index_len // n, mid_index_len % n]

            r, c = mid_index

            if matrix[r][c] < target:
                if c == n - 1:
                    s = [r+1, 0]
                else:
                    s = [r, c+1]
            elif matrix[r][c] > target:
                if c == 0:
                    e = [r-1, n - 1]
                else:
                    e = [r, c - 1]
            else:
                return True
        
        return False
        