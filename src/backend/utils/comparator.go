package utils

func Comparator(matchIdx []int, pattern string, toMatch string) bool {
	return len(matchIdx) > 0 && len(pattern) == len(toMatch)
}
