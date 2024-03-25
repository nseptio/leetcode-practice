func romanToInt(s string) int {
    romanMap := make(map[rune]int)
	romanMap['I'] = 1
	romanMap['V'] = 5
	romanMap['X'] = 10
	romanMap['L'] = 50
	romanMap['C'] = 100
	romanMap['D'] = 500
	romanMap['M'] = 1000
    
    chars := []rune(s)
    result := 0
	for i := 0; i < len(chars); i++ {
		if i == len(chars)-1 {
			result += romanMap[chars[i]]
			break
		}

		if romanMap[chars[i]] < romanMap[chars[i+1]] {
			result += romanMap[chars[i+1]] - romanMap[chars[i]]
			i++
		} else {
			result += romanMap[chars[i]]
		}
	}
    
    return result
}