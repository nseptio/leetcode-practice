import "sort"

func minimumAbsDifference(arr []int) [][]int {
    sort.Ints(arr)

    minDiff := int(^uint(0) >> 1)
    for i := 1; i < len(arr); i++ {
        minDiff = min(abs(arr[i] - arr[i-1]), minDiff)
    }

    result := make([][]int, 0)
    for i := 1; i < len(arr); i++ {
        if (abs(arr[i] - arr[i-1])) == minDiff {
            result = append(result, []int{arr[i-1], arr[i]})
        }
    }

    return result
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}