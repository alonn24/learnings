package permutations

import "log"

// rune - an alias for the type int32

// Perm - create pemutations
func Perm(input string) (result []string) {
	result = perm(input)
	return
}

func removeAtIndex(input string, i int) (selected rune, str string) {
	a := []rune(input)
	selected = a[i]
	rest := append([]rune{}, a[:i]...)
	rest = append(rest, a[i+1:]...)
	str = string(rest)
	return
}

func appendToStart(l rune, input string) string {
	x := append([]rune{l}, []rune(input)...) // convert string to rune
	return string(x)
}

func perm(input string) (res []string) {
	log.Printf("perm: %v\n", input)
	if len(input) == 1 {
		res = append(res, input)
	} else {
		for i := 0; i < len(input); i++ {
			selected, rest := removeAtIndex(input, i)
			for _, p := range perm(rest) {
				res = append(res, appendToStart(selected, p))
			}
		}
	}

	log.Printf("perm return : %v\n", res)
	return
}
