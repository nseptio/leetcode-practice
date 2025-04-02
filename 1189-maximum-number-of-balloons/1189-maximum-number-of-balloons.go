func maxNumberOfBalloons(text string) int {
	charMap := make(map[rune]int)

	for _, ch := range text {
		if string(ch) == "b" || string(ch) == "a" || string(ch) == "l" ||
			string(ch) == "o" || string(ch) == "n" {

			charMap[ch] += 1
		}
	}

	total := 0
	stillLoop := true
	for stillLoop {
		for _, ch := range "balloon" {
			if charMap[ch] <= 0 {
				stillLoop = false
				break
			}

			if string(ch) == "n" {
				total++
			}
			charMap[ch]--
		}
	}

	return total
}