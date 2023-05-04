package utils

func Comparator(matchIdx []int, pattern string, toMatch string) bool {
	return len(matchIdx) == 1 && len(pattern) == len(toMatch)
}
