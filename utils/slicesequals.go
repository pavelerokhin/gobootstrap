package utils

func SlicesEquals(s1 []float64, s2 []float64) bool {
	if len(s1) != len(s2) {
		return false
	}

	for i, v := range s2 {
		if s1[i] != v {
			return false
		}
	}

	return true
}
