package utils

func Comparator(matchIdx []int, pattern string, toMatch string) bool {
	if len(matchIdx) != 1 {
		return false
	}

	return matchIdx[0] == 0 && len(pattern) == len(toMatch)
}
