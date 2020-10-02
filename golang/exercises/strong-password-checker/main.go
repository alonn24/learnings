package checker

import (
	"errors"
	"log"
)

func isDigit(c rune) bool {
	return (c >= '0' && c <= '9')
}

func isLowerCase(c rune) bool {
	return (c >= 'a' && c <= 'z')
}

func isUpperCase(c rune) bool {
	return (c >= 'A' && c <= 'Z')
}

func analyze(s string) (int, int, error) {
	missingLowerCase := 1
	missingUpperCase := 1
	missingDigits := 1
	repeats := 0
	var err error

	repeating := 0
	chars := []rune(s)
	for i, c := range chars {
		// repeats
		if repeating == 0 {
			repeating++
		} else if chars[i-1] == c {
			repeating++
		} else {
			repeating = 1
		}
		if repeating == 3 {
			repeats++
			repeating = 0
		}

		if isDigit(c) {
			missingDigits = 0
		} else if isLowerCase(c) {
			missingLowerCase = 0
		} else if isUpperCase(c) {
			missingUpperCase = 0
		} else {
			err = errors.New("Bad input")
		}
	}
	missingSpecia := missingLowerCase + missingUpperCase + missingDigits
	return missingSpecia, repeats, err
}

const minLen = 6
const maxLen = 20

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func limitZero(a int) int {
	if a < 0 {
		return 0
	}
	return a
}

func strongPasswordChecker(s string) int {
	missingSpecia, repeats, err := analyze(s)
	insertions := limitZero(minLen - len([]rune(s)))
	deletions := limitZero(len([]rune(s)) - 20)
	if err != nil {
		log.Printf("Error: %v\n", err)
	}

	log.Printf("%v special, %v repeats,%v inserts, %v deletions", missingSpecia, repeats, insertions, deletions)

	changes := 0
	specialOverInserts := min(missingSpecia, insertions)
	changes += specialOverInserts
	missingSpecia = limitZero(missingSpecia - specialOverInserts)
	insertions = limitZero(insertions - specialOverInserts)

	specialOverRepeat := min(missingSpecia, repeats)
	changes += specialOverRepeat
	missingSpecia = limitZero(missingSpecia - specialOverRepeat)
	repeats = limitZero(repeats - specialOverRepeat)

	repeatOverDeletes := min(repeats, deletions)
	changes += repeatOverDeletes
	repeats = limitZero(repeats - repeatOverDeletes)
	deletions = limitZero(deletions - repeatOverDeletes)

	log.Printf("%v special, %v repeats,%v inserts, %v deletions, %v changes", missingSpecia, repeats, insertions, deletions, changes)
	return deletions + insertions + missingSpecia + repeats + changes
}
