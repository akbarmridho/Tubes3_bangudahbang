package utils

import "regexp"

func CleanString(input string) string {
	result := regexp.MustCompile(`\s+`).ReplaceAllString(input, " ") // replace all whitespace with single space
	result = regexp.MustCompile(`[?.;!,]`).ReplaceAllString(input, "")
	return result
}
