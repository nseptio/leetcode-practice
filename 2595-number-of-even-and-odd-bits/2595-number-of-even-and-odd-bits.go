import "strconv"

func evenOddBit(n int) []int {
	numIndicies := make([]int, 2)
	binary := strconv.FormatInt(int64(n), 2)
	for i := 0; i < len(binary); i++ {
		if string(binary[len(binary)-i-1]) == "1" {

			if i%2 == 0 {
				numIndicies[0]++
			} else {
				numIndicies[1]++
			}
		}
	}

	return numIndicies
}