func maxArea(height []int) int {
    left := 0
    right := len(height) - 1
    maxArea := -1
    for _ = range len(height) {
        h := min(height[left], height[right])
        w := right - left
        area := h * w
        if area > maxArea {
            maxArea = area
        }

        if AbsInt(left - right) != 1 {
            if height[left] < height[right] {
                left++
            } else {
                right--
            }
        }
    }

    return maxArea
}


func AbsInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}