package utils

import (
	// "fmt"
	"strings"
)

// LevenshteinDistance Calculate Levenshtein Distance of two string (how many operations needed to match the string)
func LevenshteinDistance(input string, toMatch string) int {
	str1 := strings.ToLower(input)
	str2 := strings.ToLower(toMatch)
	len1 := len(str1)
	len2 := len(str2)

	matrix := make([][]int, len1+1)
	for i := 0; i <= len1; i++ {
		matrix[i] = make([]int, len2+1)
		matrix[i][0] = i
	}
	for j := 0; j <= len2; j++ {
		matrix[0][j] = j
	}

	for j := 1; j <= len2; j++ {
		for i := 1; i <= len1; i++ {
			if str1[i-1] == str2[j-1] {
				matrix[i][j] = matrix[i-1][j-1]
			} else {
				minimum := matrix[i-1][j]
				if matrix[i][j-1] < minimum {
					minimum = matrix[i][j-1]
				}
				if matrix[i-1][j-1] < minimum {
					minimum = matrix[i-1][j-1]
				}
				matrix[i][j] = minimum + 1
			}
		}
	}

	return matrix[len1][len2]
}

// MeasureSimilarity Calculate similarity in percent using Levenshtein Distance
func MeasureSimilarity(input string, toMatch string) float64 {
	distance := LevenshteinDistance(input, toMatch)
	maxLen := len(input)
	if len(toMatch) > maxLen {
		maxLen = len(toMatch)
	}
	similarity := 1 - (float64(distance) / float64(maxLen))
	return similarity
}

// Testing
// func main() {
// 	str1 := "hello this is livia"
// 	str2 := "hella this is livia"
// 	similarity := MeasureSimilarity(str1, str2)
// 	fmt.Printf("Similarity between '%s' and '%s' is %.2f%%\n", str1, str2, similarity)

// 	str3 := "what is the capital city of australia"
// 	str4 := "what is th ecapthial city of austalisa"
// 	similarity = MeasureSimilarity(str3, str4)
// 	fmt.Printf("Similarity between '%s' and '%s' is %.2f%%\n", str3, str4, similarity)
// }
