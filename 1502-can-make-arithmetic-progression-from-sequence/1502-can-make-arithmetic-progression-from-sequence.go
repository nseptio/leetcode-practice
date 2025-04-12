import "sort"

func canMakeArithmeticProgression(arr []int) bool {
    sort.Ints(arr)
    diff := 0
    for i := 1; i < len(arr); i++ {
        if i == 1 {
            diff = abs(arr[i-1] - arr[i])
            continue
        }

        if diff != abs(arr[i-1] - arr[i]) {
            return false
        }
    }
    return true
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}