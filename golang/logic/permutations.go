package permutations

// rune - an alias for the type int32

// Perm - create pemutations
func Perm(input string) (result [][]rune) {
	r := []rune(input)
	result = perm(r, 0)
	return
}

func perm(a []rune, i int) (res [][]rune) {
	if i > len(a) {
		return append(res, a)
	}
	res = append(res, perm(a, i+1)...)
	for j := i + 1; j < len(a); j++ {
		a[i], a[j] = a[j], a[i]
		res = append(res, perm(a, i+1)...)
		a[i], a[j] = a[j], a[i]
	}
	return
}
