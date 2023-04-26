package string_matcher

import (
	// "fmt"
	"strings"
)

// import (
// 	"fmt"
// )

func BadCharHeuristic(pattern string, patLen int) []int {
	ret := make([]int, 0)

	for i := 0; i < 256; i++ {
		ret = append(ret, -1)
	}
	for i := 0; i < patLen; i++ {
		ret[int(pattern[i])] = i
	}
	return ret
}

func BM(patternStr string, toMatchStr string) []int {
	// todo
	patternStr = strings.ToLower(patternStr)
	toMatchStr = strings.ToLower(toMatchStr)
	var patLen int = len(patternStr)
	var textLen int = len(toMatchStr)
	ret := make([]int, 0)
	if patLen <= textLen {
		var badChar []int = BadCharHeuristic(patternStr, patLen)
		i := 0
		for i <= textLen-patLen {
			j := patLen - 1
			for j >= 0 && patternStr[j] == toMatchStr[i+j] {
				j--
			}

			if j < 0 { // if it is a matching string
				ret = append(ret, i)
				if i+patLen < textLen {
					i += patLen - badChar[toMatchStr[i+patLen]]
				} else {
					i += 1
				}
			} else {
				if diff := j - badChar[toMatchStr[i+j]]; diff > 1 {
					i += diff
				} else {
					i += 1
				}
			}
		}
	}
	return ret
}

// func main() {
// 	txt := "saya suka makan nasi? apakah kamu juga suka makan nasi? entahlah ayok makaan nasi"
// 	pattern := "MakAn naSi"
// 	var res []int = BM(pattern, txt)
// 	fmt.Println(res)
// }
