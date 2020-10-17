package atoi

import (
	"math"
)

func isIn(c rune, arr []rune) bool {
	for _, r := range arr {
		if r == c {
			return true
		}
	}
	return false
}

func takeWhile(s string, c []rune) (string, string) {
	chars := []rune(s)
	for i := 0; i < len(chars); i++ {
		if !isIn(chars[i], c) {
			return string(chars[0:i]), string(chars[i:len(chars)])
		}
	}
	return s, ""
}

func getPower(i int) float64 {
	power := math.Pow(10, float64(i-1))
	return power
}

const maxUint = ^uint32(0)
const maxInt = float64(maxUint >> 1)
const minInt = (float64(maxUint>>1) + 1) * -1

func myAtoi(s string) int {
	_, withoutSpaces := takeWhile(s, []rune{' '})
	plusOrMinus, rest := takeWhile(withoutSpaces, []rune{'-', '+'})
	if len(plusOrMinus) > 1 {
		return 0
	}
	number, _ := takeWhile(rest, []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'})

	res := float64(0)

	for i, n := range []rune(number) {
		power := getPower(len(number) - i)
		res = res + float64(n-'0')*power
	}

	if plusOrMinus == "-" {
		res = res * (-1)
	}
	if res > maxInt {
		res = maxInt
	}
	if res < minInt {
		res = minInt
	}
	return int(res)
}
