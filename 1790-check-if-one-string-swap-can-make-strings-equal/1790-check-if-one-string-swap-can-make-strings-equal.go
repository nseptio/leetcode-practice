func areAlmostEqual(s1 string, s2 string) bool {
	s1Map := make(map[byte]int)
	s2Map := make(map[byte]int)
	for i := 0; i < len(s1); i++ {
		s1Map[s1[i]]++
		s2Map[s2[i]]++
	}

	for k, v := range s1Map {
		if s2Map[k] != v {
			return false
		}
	}

	diff := 0
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			diff++
		}
	}

	if diff > 2 {
		return false
	}

	return true
}