package checker

import (
	"log"
)

func isLower(c byte) bool {
	return c >= 'a' && c <= 'z'
}

func isUpper(c byte) bool {
	return c >= 'A' && c <= 'Z'
}

func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func analyze(s string) (int, []int) {
	missingLower := 1
	missingUpper := 1
	missingDigit := 1
	repeats := make([]int, len(s))

	for i := 0; i < len(s); {
		c := s[i]
		if isLower(c) {
			missingLower = 0
		} else if isUpper(c) {
			missingUpper = 0
		} else if isDigit(c) {
			missingDigit = 0
		}

		f := i
		for i < len(s) && s[i] == c {
			i++
		}
		repeats[f] = i - f
	}
	return missingLower + missingUpper + missingDigit, repeats
}

const minLength = 6
const maxLength = 20

func strongPasswordChecker(s string) int {
	missingSpecial, repeats := analyze(s)

	// Insertions
	if len(s) < 6 {
		return max(missingSpecial, minLength-len(s))
	}
	// Deletions
	deletions := max(len(s)-maxLength, 0)
	log.Printf("%v deletions", deletions)

	minChanges := deletions

	// remove repeats by deletions
	log.Printf("%v", s)
	log.Printf("0 - %v", repeats)
	for r := 1; r < 3; r++ {
		for i := 0; i < len(repeats) && deletions > 0; i++ {
			if repeats[i] >= 3 && repeats[i]%3 == (r-1) {
				repeats[i] -= min(deletions, r)
				deletions -= r
			}
		}
		log.Printf("%v - %v", r, repeats)
	}

	// count replacing
	replaces := 0
	for i := 0; i < len(repeats); i++ {
		if repeats[i] >= 3 && deletions > 0 {
			temp := repeats[i] - 2
			repeats[i] -= deletions
			deletions -= temp
		}

		if repeats[i] >= 3 {
			replaces += repeats[i] / 3
		}
	}

	return minChanges + max(missingSpecial, replaces)
}
