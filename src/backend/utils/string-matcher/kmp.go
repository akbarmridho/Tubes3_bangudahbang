package string_matcher

import (
	"strings"
)

func KMP(patternStr string, toMatchStr string) []int {
	patternStr = strings.ToLower(patternStr)
	toMatchStr = strings.ToLower(toMatchStr)
	result := make([]int, 0)

	// use rune to handle UTF-8 chars instead of ascii
	pattern := []rune(patternStr)
	toMatch := []rune(toMatchStr)
	patternLen := len(pattern)

	// compute border
	borders := make([]int, patternLen)
	j, i := 0, 1

	for i < patternLen {
		if pattern[j] == pattern[i] {
			// j+1 chars match
			borders[i] = j + 1
			i++
			j++
		} else if j > 0 { // j follows matching prefix
			j = borders[j-1]
		} else { // no match
			borders[i] = 0
			i++
		}
	}

	// begin matching
	i, j = 0, 0

	for i < len(toMatch) {
		if pattern[j] == toMatch[i] {
			if j == patternLen-1 {
				result = append(result, i-patternLen+1) // found match
			}
			i++
			j++
		} else if j > 0 {
			j = borders[j-1]
		} else {
			i++
		}
	}

	return result
}
