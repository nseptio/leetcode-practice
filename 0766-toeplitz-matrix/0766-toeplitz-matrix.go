func isToeplitzMatrix(matrix [][]int) bool {
    var rowBefore []int
    for i, row := range matrix {
        if i == 0 {
            continue
        }
        rowBefore = matrix[i-1]
        for j, col := range row {
            if j == 0 {
                continue
            }
            if col != rowBefore[j-1] {
                return false
            }
        }
    }
    
    return true
}